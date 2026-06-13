package persist

import (
	"bytes"
	"encoding/json"
	"errors"
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/javahongxi/golab/crawler/engine"
)

func ItemMockSaver() chan engine.Item {
	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++
		}
	}()

	return out
}

func ItemSaver(
	index string) (chan engine.Item, error) {
	cfg := elasticsearch.Config{
		Addresses: []string{"http://localhost:9200"},
	}
	client, err := elasticsearch.NewClient(cfg)

	if err != nil {
		return nil, err
	}

	out := make(chan engine.Item)
	go func() {
		itemCount := 0
		for {
			item := <-out
			log.Printf("Item Saver: got item #%d: %v", itemCount, item)
			itemCount++

			err := Save(client, index, item)
			if err != nil {
				log.Printf("Item Saver: error saving item %v: %v", item, err)
			}
		}
	}()

	return out, nil
}

func Save(client *elasticsearch.Client, index string,
	item engine.Item) error {

	if item.Type == "" {
		return errors.New("must supply Type")
	}

	data, err := json.Marshal(item)
	if err != nil {
		return err
	}

	if item.Id != "" {
		_, err = client.Index(index, bytes.NewReader(data), client.Index.WithDocumentID(item.Id))
	} else {
		_, err = client.Index(index, bytes.NewReader(data))
	}

	return err
}
