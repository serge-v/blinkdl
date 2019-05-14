package main

import (
	"flag"
	"fmt"
	"log"
	"os"

	"github.com/serge-v/autocomplete"
	"golang.org/x/crypto/ssh/terminal"
)

var (
	ac           = flag.Bool("c", false, "for autocomplete support")
	login        = flag.String("login", "", "login using `EMAIL`. Put password into XPASSWD env var")
	list         = flag.Bool("list", false, "list videos")
	listTemplate = flag.String("list-template", "", "item `template` for -list command. View all fields: '{{printf \"%+v\\n\\n\"  .}}'")
	download     = flag.Bool("download", false, "download all videos into ~/.local/blink")
	test         = flag.Bool("test", false, "do test")
	dryRun       = flag.Bool("dry", false, "use cached data")
	daysBack     = flag.Int("days", 1, "`DAYS` back")
	page         = flag.Int("page", 1, "page number `NUM`")
	cameraConfig = flag.String("camera-info", "", "get camera info by `NAME`")
	info         = flag.Bool("info", false, "get system info")
)

var (
	autocompleteScript = `complete -C "blinkdl -c" blinkdl`
	acfile             = os.Getenv("HOME") + "/.config/bash_completion/blinkdl"
)

func main() {
	log.SetFlags(log.LstdFlags | log.Lshortfile)
	flag.Parse()
	autocomplete.Init(acfile, autocompleteScript)
	if *ac {
		autocomplete.Print()
		return
	}

	var err error
	cli := NewClient()
	switch {
	case *login != "":
		pwd := os.Getenv("XPASSWD")
		if pwd == "" {
			fmt.Println("password?")
			buf, err := terminal.ReadPassword(0)
			if err != nil {
				log.Fatal(err)
			}
			pwd = string(buf)
		}
		err = cli.Login(*login, pwd)
	case *test:
		err = cli.doTest()
	case *info:
		err = cli.PrintSystemInfo()
	case *list:
		err = cli.List(*listTemplate, *daysBack, *page)
	case *download:
		err = cli.Download(*daysBack, *page)
	case *cameraConfig != "":
		err = cli.GetCameraConfig(*cameraConfig)
	}

	if err != nil {
		log.Fatalf("%+v", err)
	}

}
