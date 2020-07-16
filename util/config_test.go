package util

import (
	"testing"

	"github.com/spf13/viper"
)


func TestConfig(t *testing.T) {
	err:= InitConfig("../conf/config.yaml")
	if err!= nil {
		t.Error(err)
	}

	t.Log(viper.GetString("frontend.port"))
	t.Log(viper.GetString("elastic.host"))
	t.Log(viper.GetString("spider.worker_count"))
}