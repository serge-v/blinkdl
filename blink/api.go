// Package blink provides some API to download videos from Blink security home system.
package blink

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"strconv"
	"strings"
	"text/template"
	"time"

	"github.com/pkg/errors"
)

const (
	host            = "prod.immedia-semi.com"
	loginURL        = "https://rest.prod.immedia-semi.com/login"
	timestampFormat = "2006-01-02T15:04:05-0700"
	regionalHost    = "https://rest-%s.immedia-semi.com"
)

type login struct {
	ClientSpecifier string `json:"client_specifier"`
	Email           string `json:"email"`
	Password        string `json:"password"`
}

type network struct {
	Name      string
	Onboarded bool
}

type loginResponse struct {
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
	Networks map[string]network
	Region   map[string]string
}

// Client is blink API client.
type Client struct {
	token     string
	region    string
	accountID int64
	networkID int64
	cli       *http.Client
	cacheDir  string
	localDir  string
	DryRun    bool // get responses from cache
	Debug     bool
}

// NewClient creates new instance of API client.
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

	var a loginResponse
	if err := json.NewDecoder(f).Decode(&a); err != nil {
		log.Fatal(err)
	}

	c.token = a.Authtoken.Authtoken
	for k := range a.Region {
		c.region = k
	}
	c.accountID = a.Account.ID
	for n := range a.Networks {
		c.networkID, _ = strconv.ParseInt(n, 10, 64)
	}
}

// Login calls /login endpoint and receives authentication info.
func (c *Client) Login(email, password string) error {
	if password == "" {
		return fmt.Errorf("empty password")
	}
	login := login{
		Email:           email,
		Password:        password,
		ClientSpecifier: "iPhone XR | 3.4.12 | 349300",
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

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("invalid status: %s", resp.Status)
	}

	var loginResp loginResponse
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
	if c.DryRun {
		println("=== dry", url, name)
		return ioutil.ReadFile(c.cacheDir + "/" + name)
	}

	var u string
	if strings.HasPrefix(url, "https://") {
		u = url
	} else {
		u = fmt.Sprintf(regionalHost, c.region) + url
	}

	if c.Debug {
		log.Println("GET", u)
	}
	req, err := http.NewRequest(http.MethodGet, u, nil)
	if err != nil {
		return nil, errors.Wrap(err, "new request")
	}

	req.Header.Set("TOKEN_AUTH", c.token)
	req.Header.Set("Host", host)

	resp, err := c.cli.Do(req)
	if err != nil {
		return nil, errors.Wrap(err, "do request")
	}

	defer resp.Body.Close()
	buf, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, errors.Wrap(err, "read body")
	}
	if c.Debug {
		log.Println("resp:", string(buf))
	}

	if name != "" {
		var b bytes.Buffer
		if err := json.Indent(&b, buf, "", "\t"); err != nil {
			log.Println(string(buf))
			return nil, errors.Wrap(err, "reindent")
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

func (c *Client) getVideos(daysBack, page int) (*List, error) {
	ts := time.Now().UTC().Add(-time.Hour * 24 * time.Duration(daysBack)).Format(timestampFormat)
	//	path := fmt.Sprintf("/api/v2/videos/changed?since=%s&page=%d", ts, page)
	path := fmt.Sprintf("/api/v1/accounts/%d/media/changed?since=%s&page=%d", c.accountID, ts, page)
	//	path := fmt.Sprintf("/api/v2/videos/page/%d", page)

	buf, err := c.request(path, "list.json")
	if err != nil {
		return nil, errors.Wrapf(err, "url: %s", path)
	}

	var list List
	if err := json.Unmarshal(buf, &list); err != nil {
		return nil, err
	}

	return &list, nil
}

// List prints recent video list.
func (c *Client) List(tmpl string, daysBack, page int) error {
	var err error
	var t *template.Template

	if tmpl != "" {
		t, err = template.New("item").Parse(tmpl)
		if err != nil {
			return errors.Wrap(err, "template")
		}
	}

	list, err := c.getVideos(daysBack, page)
	if err != nil {
		return errors.Wrap(err, "getVideos")
	}

	for _, e := range list.Media {
		if t != nil {
			if err := t.Execute(os.Stdout, e); err != nil {
				return err
			}
		} else {
			viewed := "*"
			if e.Watched {
				viewed = " "
			}

			fmt.Println(e.CreatedAt, viewed, e.DeviceName)
		}
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

// Download downloads recent videos into ~/.local/blink directory.
func (c *Client) Download(pages int) error {
	timestamp := "1970-01-01T00:00Z"

	//	timestampFormat = "2006-01-02T15:04:05-0700"
	for page := 1; page < pages; page++ {
		log.Println("page", page)
		path := fmt.Sprintf("/api/v1/accounts/%d/media/changed?since=%s&page=%d", c.accountID, timestamp, page)

		buf, err := c.request(path, "download.json")
		if err != nil {
			return errors.Wrapf(err, "url: %s", path)
		}

		var list List
		if err := json.Unmarshal(buf, &list); err != nil {
			return err
		}

		err = c.downloadPage(list)
		if err != nil {
			return err
		}
	}
	return nil
}

func (c *Client) downloadPage(list List) error {
	for _, e := range list.Media {
		if e.Deleted {
			continue
		}
		ts, err := time.Parse("2006-01-02T15:04:05-07:00", e.CreatedAt)
		if err != nil {
			return err
		}

		fname := c.localDir + ts.Format("20060102-150405-") + e.DeviceName + ".mp4"
		fi, err := os.Stat(fname)
		if err == nil && fi.Size() > 0 {
			continue
		}

		buf, err := c.request(e.Media, "")
		if err != nil {
			return err
		}
		if err := ioutil.WriteFile(fname, buf, 0700); err != nil {
			return err
		}
		log.Println(fname, e.UpdatedAt)
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

func (c *Client) getHomeScreen() (*Homescreen, error) {
	const path = "/api/v3/accounts/%d/homescreen"

	buf, err := c.request(fmt.Sprintf(path, c.accountID), "homescreen.json")
	if err != nil {
		return nil, err
	}

	var hs Homescreen
	if err := json.Unmarshal(buf, &hs); err != nil {
		return nil, err
	}
	return &hs, nil
}

// PrintSystemInfo prints some useful info.
func (c *Client) PrintSystemInfo() error {
	hs, err := c.getHomeScreen()
	if err != nil {
		return err
	}

	fmt.Println("SyncModules")
	for i, sm := range hs.SyncModules {
		fmt.Println("   ", i+1, sm.Name, "status:", sm.Status, "wifi:", sm.WifiStrength)
	}
	var cameras []string
	fmt.Println("Cameras")
	for i, c := range hs.Cameras {
		fmt.Println("   ", i+1, c.Name, "battery:", c.Battery)
		cameras = append(cameras, c.Name)
	}
	fmt.Println("Storage")

	fmt.Println("    Usage:", hs.VideoStats.Storage, "%")
	fmt.Println("    AutoDelete:", hs.VideoStats.AutoDeleteDays, "days")
	ioutil.WriteFile(os.Getenv("HOME")+"/.cache/blink/cameras.txt", []byte(strings.Join(cameras, "\n")), 0600)

	return nil
}

// GetCameraConfig prints camera info.
func (c *Client) GetCameraConfig(name string) error {
	hs, err := c.getHomeScreen()
	if err != nil {
		return err
	}

	var camID int64
	for _, c := range hs.Cameras {
		if c.Name == name {
			camID = c.ID
			break
		}
	}

	if camID == 0 {
		fmt.Println("camera not found")
	}

	path := fmt.Sprintf("/network/%d/camera/%d/config", c.networkID, camID)
	buf, err := c.request(path, fmt.Sprintf("camera-%d.json", camID))
	if err != nil {
		return err
	}

	var cc CameraConfig
	if err := json.Unmarshal(buf, &cc); err != nil {
		return err
	}

	cam := cc.Camera[0]
	fmt.Println("Name:", cam.Name)
	fmt.Printf("AlertInterval: %ds\n", cam.AlertInterval)
	fmt.Printf("VideoLength: %ds\n", cam.VideoLength)
	fmt.Printf("Temperature: %dF\n", cam.Temperature)
	fmt.Printf("BatteryVoltage: %.2fV\n", float64(cam.BatteryVoltage)/100.0)

	return nil
}

// DoTest is a place for testing new methods.
func (c *Client) DoTest() error {
	return nil
}
