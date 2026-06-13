package persist

import (
	"log"

	"github.com/elastic/go-elasticsearch/v8"
	"github.com/javahongxi/golab/crawler/engine"
	"github.com/javahongxi/golab/crawler/persist"
)

type ItemSaverService struct {
	Client *elasticsearch.Client
	Index  string
}

func (s *ItemSaverService) Save(
	item engine.Item, result *string) error {
	err := persist.Save(s.Client, s.Index, item)
	log.Printf("Item %v saved.", item)
	if err == nil {
		*result = "ok"
	} else {
		log.Printf("Error saving item %v: %v", item, err)
	}
	return err
}
