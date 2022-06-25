package main

import (
	"linebot/config"
	"linebot/internal/route"

	_ "github.com/lib/pq"
)

func main() {

	ginroute := route.InitRouter()

	ginroute.Run(":" + config.HttpPort)
}
