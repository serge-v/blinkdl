package main

import (
	"flag"
	"log"
	"os"
)

var (
	login        = flag.String("login", "", "login using `EMAIL`. Put password into XPASSWD env var")
	list         = flag.Bool("list", false, "list videos")
	listTemplate = flag.String("list-template", "", "item `template` for -list command. View all fields: '{{printf \"%+v\\n\\n\"  .}}'")
	download     = flag.Bool("download", false, "download all videos into ~/.local/blink")
	test         = flag.Bool("test", false, "do test")
	dryRun       = flag.Bool("dry", false, "use cached data")
)

func main() {
	flag.Parse()
	var err error
	cli := NewClient()
	switch {
	case *login != "":
		pwd := os.Getenv("XPASSWD")
		err = cli.Login(*login, pwd)
	case *test:
		err = cli.getHomeScreen()
	case *list:
		err = cli.List(*listTemplate)
	case *download:
		err = cli.Download()
	}

	if err != nil {
		log.Fatal(err)
	}

}
