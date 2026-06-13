package main

import (
	"flag"
	"fmt"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/javahongxi/golab/crawler/config"
	"github.com/javahongxi/golab/crawler_distributed/persist"
	"github.com/javahongxi/golab/crawler_distributed/rpcsupport"
)

var port = flag.Int("port", 0, "the port for me to listen on")

func main() {
	flag.Parse()
	if *port == 0 {
		fmt.Println("must specify a port")
		return
	}
	log.Fatal(serveRpc(fmt.Sprintf(":%d", *port), config.ElasticIndex))
}

func serveRpc(host, index string) error {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	client, err := elasticsearch.NewClient(cfg)
	if err != nil {
		return err
	}

	return rpcsupport.ServeRpc(host,
		&persist.ItemSaverService{
			Client: client,
			Index:  index,
		})
}
