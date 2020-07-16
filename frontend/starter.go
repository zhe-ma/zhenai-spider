package main

import (
	"fmt"
	"net/http"
	"zhenai-spider/frontend/controller"
	"zhenai-spider/util"

	"github.com/spf13/viper"
)

func main() {
	err := util.InitConfig("../conf/config.yaml")
	if err != nil {
		panic(err)
	}

	http.Handle("/res/", http.StripPrefix("/res/", http.FileServer(http.Dir("./view/res"))))
	http.Handle("/", controller.NewSearchResultHandler("view/index.html", viper.GetString("elastic.host")))

	err = http.ListenAndServe(viper.GetString("frontend.port"), nil)
	if err != nil {
		fmt.Println("Failed to start HTTP server, error:", err)
	}
}
