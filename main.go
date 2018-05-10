package main

import (
	"flag"
	"os"

	"github.com/xkon/hotspot/server"
	"github.com/xkon/hotspot/spider"
)

func main() {
	mode := flag.String("m", "spider", "run in server/spider mode")
	flag.Parse()
	if flag.NFlag() < 1 {
		flag.Usage()
		os.Exit(0)
	}
	switch *mode {
	case "server":
		server.ListenAndServe()
	case "spider":
		spider.Run()
	default:
		spider.Run()
	}
}
