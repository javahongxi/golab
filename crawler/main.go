package main

import (
	"github.com/javahongxi/golab/crawler/config"
	"github.com/javahongxi/golab/crawler/engine"
	"github.com/javahongxi/golab/crawler/persist"
	"github.com/javahongxi/golab/crawler/scheduler"
	"github.com/javahongxi/golab/crawler/zhenai/parser"
)

func main() {
	itemChan := persist.ItemMockSaver()

	e := engine.ConcurrentEngine{
		Scheduler:        &scheduler.QueuedScheduler{},
		WorkerCount:      100,
		ItemChan:         itemChan,
		RequestProcessor: engine.Worker,
	}

	e.Run(engine.Request{
		Url:    "http://localhost:8080/mock/www.zhenai.com/zhenghun",
		Parser: engine.NewFuncParser(parser.ParseCityList, config.ParseCityList),
	})
}
