package main

import (
	"linebot/config"

	"linebot/internal/route"

	_ "github.com/lib/pq"
)

func main() {
	//model.InitDbContext()
	//db.InitDbContext()

	ginroute := route.InitRouter()

	ginroute.Run(":" + config.HttpPort)

	//first page
}
