package db

import "github.com/astaxie/beego"

var maxPool int

func init() {
	var err error
	maxPool, err = beego.AppConfig.Int("DBMaxPool")
	if err != nil {
		//todo:panic!!
		println(err)
	}

	checkAndInitServiceConnection()
}

func checkAndInitServiceConnection() {
	if service.baseSession == nil {
		service.URL = beego.AppConfig.String("DBPath")
		err := service.New()

		if err != nil {
			//todo: panic!!
			println(err)
		}
	}
}
