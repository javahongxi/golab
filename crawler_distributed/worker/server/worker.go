package main

import (
	"fmt"

	"log"

	"flag"

	"github.com/javahongxi/golab/crawler/fetcher"
	"github.com/javahongxi/golab/crawler_distributed/rpcsupport"
	"github.com/javahongxi/golab/crawler_distributed/worker"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	fetcher.SetVerboseLogging()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(rpcsupport.ServeRpc(fmt.Sprintf(":%d", *port), worker.CrawlService{}))
}
