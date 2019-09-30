package main

import (
	"github.com/javahongxi/whatsgo/crawler/config"
	"github.com/javahongxi/whatsgo/crawler/engine"
	"github.com/javahongxi/whatsgo/crawler/persist"
	"github.com/javahongxi/whatsgo/crawler/scheduler"
	"github.com/javahongxi/whatsgo/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(
		config.ElasticIndex)
	if err != nil {
		panic(err)
	}

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url: "http://www.starter.url.here",
		Parser: engine.NewFuncParser(
			parser.ParseCityList,
			config.ParseCityList),
	})
}
