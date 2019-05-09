package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"text/template"
	"time"
)

const (
	host            = "prod.immedia-semi.com"
	loginURL        = "https://rest.prod.immedia-semi.com/login"
	timestampFormat = "2006-01-02T15:04:05-0700"
	regionalHost    = "https://rest-%s.immedia-semi.com"
)

type Login struct {
	ClientSpecifier string `json:"client_specifier"`
	Email           string `json:"email"`
	Password        string `json:"password"`
}

type Network struct {
	Name      string
	Onboarded bool
}

type LoginResponse struct {
	Account struct {
		ID int64
	}
	Authtoken struct {
		Authtoken string
		Message   string
	}
	Client struct {
		ID int64
	}
	Networks map[string]Network
	Region   map[string]string
}

type List struct {
	Limit   int64 `json:"limit"`
	PurgeID int64 `json:"purge_id"`
	Videos  []struct {
		AccountID       int64       `json:"account_id"`
		Address         string      `json:"address"`
		CameraID        int64       `json:"camera_id"`
		CameraName      string      `json:"camera_name"`
		CreatedAt       string      `json:"created_at"`
		Deleted         bool        `json:"deleted"`
		Description     string      `json:"description"`
		Encryption      string      `json:"encryption"`
		EncryptionKey   interface{} `json:"encryption_key"`
		EventID         interface{} `json:"event_id"`
		ID              int64       `json:"id"`
		Length          int64       `json:"length"`
		Locked          bool        `json:"locked"`
		NetworkID       int64       `json:"network_id"`
		NetworkName     string      `json:"network_name"`
		Partial         bool        `json:"partial"`
		Ready           bool        `json:"ready"`
		Size            int64       `json:"size"`
		StorageLocation string      `json:"storage_location"`
		Thumbnail       string      `json:"thumbnail"`
		TimeZone        string      `json:"time_zone"`
		UpdatedAt       string      `json:"updated_at"`
		UploadTime      int64       `json:"upload_time"`
		Viewed          string      `json:"viewed"`
	} `json:"videos"`
}

type Client struct {
	token     string
	region    string
	accountID int64
	cli       *http.Client
	cacheDir  string
	localDir  string
}

func NewClient() *Client {
	c := &Client{
		cli:      &http.Client{},
		cacheDir: os.Getenv("HOME") + "/.cache/blink/",
		localDir: os.Getenv("HOME") + "/.local/blink/",
	}
	os.MkdirAll(c.cacheDir, 0700)
	os.MkdirAll(c.localDir, 0700)
	c.loadAuth()
	return c
}

func (c *Client) loadAuth() {
	f, err := os.Open(c.cacheDir + "/auth.json")
	if os.IsNotExist(err) {
		return
	}

	if err != nil {
		log.Fatal(err)
	}

	var a LoginResponse
	if err := json.NewDecoder(f).Decode(&a); err != nil {
		log.Fatal(err)
	}

	c.token = a.Authtoken.Authtoken
	for k := range a.Region {
		c.region = k
	}
	c.accountID = a.Account.ID
}

func (c *Client) Login(email, password string) error {
	if password == "" {
		return fmt.Errorf("empty password")
	}
	login := Login{
		Email:           email,
		Password:        password,
		ClientSpecifier: "iPhone 9.2 | 2.2 | 222",
	}

	var body bytes.Buffer
	if err := json.NewEncoder(&body).Encode(login); err != nil {
		return err
	}

	req, err := http.NewRequest(http.MethodPost, loginURL, &body)
	if err != nil {
		return err
	}

	req.Header.Set("Content-Type", "application/json")
	req.Header.Set("Host", host)

	resp, err := c.cli.Do(req)
	if err != nil {
		return err
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return err
	}
	log.Printf("login resp: %s", string(buf))

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status: %s", resp.Status)
	}

	var loginResp LoginResponse
	if err := json.Unmarshal(buf, &loginResp); err != nil {
		return err
	}
	if loginResp.Authtoken.Authtoken == "" {
		return fmt.Errorf("login failed")
	}

	var b bytes.Buffer
	if err := json.Indent(&b, buf, "", "\t"); err != nil {
		return err
	}

	if err := ioutil.WriteFile(c.cacheDir+"/auth.json", b.Bytes(), 0600); err != nil {
		return err
	}
	c.loadAuth()
	return nil
}

func (c *Client) request(url, name string) ([]byte, error) {
	if *dryRun {
		println("=== dry", name)
		return ioutil.ReadFile(c.cacheDir + "/" + name)
	}

	u := fmt.Sprintf(regionalHost, c.region) + url
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Set("TOKEN_AUTH", c.token)
	req.Header.Set("Host", host)

	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, err
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	if name != "" {
		var b bytes.Buffer
		if err := json.Indent(&b, buf, "", "\t"); err != nil {
			return nil, err
		}
		if err := ioutil.WriteFile(c.cacheDir+"/"+name, b.Bytes(), 0600); err != nil {
			return nil, err
		}
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("invalid status: %s", resp.Status)
	}

	return buf, nil
}

func (c *Client) getVideos() (*List, error) {
	const path = "/api/v2/videos/changed?since=%s&page=1"

	ts := time.Now().UTC().Add(-time.Hour * 24).Format(timestampFormat)
	buf, err := c.request(fmt.Sprintf(path, ts), "list.json")
	if err != nil {
		return nil, err
	}

	var list List
	if err := json.Unmarshal(buf, &list); err != nil {
		return nil, err
	}

	return &list, nil
}

func (c *Client) List(tmpl string) error {
	var err error
	var t *template.Template

	if tmpl != "" {
		t, err = template.New("item").Parse(tmpl)
		if err != nil {
			return err
		}
	}

	list, err := c.getVideos()
	if err != nil {
		return err
	}

	for _, e := range list.Videos {
		if t != nil {
			if err := t.Execute(os.Stdout, e); err != nil {
				return err
			}
		} else {
			viewed := " "
			if e.Viewed == "" {
				viewed = "*"
			}

			fmt.Println(e.CreatedAt, viewed, e.CameraName, e.Length)
		}
	}
	return nil
}

func (c *Client) Download() error {
	list, err := c.getVideos()
	if err != nil {
		return err
	}

	for _, e := range list.Videos {
		if e.Deleted {
			continue
		}
		ts, err := time.Parse("2006-01-02T15:04:05-07:00", e.CreatedAt)
		if err != nil {
			return err
		}

		fname := c.localDir + ts.Format("20060102-150405-") + e.CameraName + ".mp4"
		fi, err := os.Stat(fname)
		if err == nil && fi.Size() > 0 {
			continue
		}

		buf, err := c.request(e.Address, "")
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(fname, buf, 0700); err != nil {
			return err
		}
		log.Println(fname, e.Size)
	}

	return nil
}

func (c *Client) getHomeScreen() error {
	const path = "/api/v3/accounts/%d/homescreen"

	_, err := c.request(fmt.Sprintf(path, c.accountID), "homescreen.json")
	if err != nil {
		return err
	}

	return nil
}
