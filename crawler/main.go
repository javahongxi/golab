package main

import (
	"golab/crawler/config"
	"golab/crawler/engine"
	"golab/crawler/persist"
	"golab/crawler/scheduler"
	"golab/crawler/zhenai/parser"
)

func main() {
	itemChan, err := persist.ItemSaver(config.ElasticIndex)
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
		Url:    "http://www.starter.url.here",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
