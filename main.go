package main

import (
	"ginchat/models"
	"ginchat/router"
	"ginchat/utils"
	"time"

	"github.com/spf13/viper"
)

// test
func main() {
	//init midware
	utils.InitConfig()
	utils.InitMySQL()
	utils.InitRedis()
	InitTimer()
	//create server
	r := router.Router()
	//run the server in specific port
	r.Run(viper.GetString("port.server"))
}

// 初始化定时器
func InitTimer() {
	utils.Timer(time.Duration(viper.GetInt("timeout.DelayHeartbeat"))*time.Second, time.Duration(viper.GetInt("timeout.HeartbeatHz"))*time.Second, models.CleanConnection, "")
}
