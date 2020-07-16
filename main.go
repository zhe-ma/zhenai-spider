package main

import (
	"zhenai-spider/database"
	"zhenai-spider/engine"
	"zhenai-spider/fetcher"
	"zhenai-spider/scheduler"
	"zhenai-spider/util"

	"github.com/spf13/viper"
)

func main() {
	util.InitLogging()
	err := util.InitConfig("conf/config.yaml")
	if err != nil {
		util.ErrorLog.Fatalln(err)
	}

	dataSaver := database.NewElasticSaver(viper.GetString("elastic.host"))
	// simpleEngine := engine.SimpleEngine{DataSaver: dataSaver}
	// simpleEngine.Run(&fetcher.CityFetcher{URL: "http://www.zhenai.com/zhenghun"})

	// scheduler := &scheduler.SimpleScheduler{}
	scheduler := &scheduler.QueuedScheduler{}

	concurrentEngine := engine.ConcurrentEngine{WorkerCount: viper.GetInt("spider.woker_count"), Scheduler: scheduler, DataSaver: dataSaver}
	concurrentEngine.Run(&fetcher.CityFetcher{URL: "http://www.zhenai.com/zhenghun"})
}
