package main

import (
	"linebot/config"

	"linebot/internal/repository"
	"linebot/internal/route"

	_ "github.com/lib/pq"
)

func main() {
	//repository.InitDbContext()
	repository.InitDbContext()

	ginroute := route.InitRouter()

	ginroute.Run(":" + config.HttpPort)

	//first page
}
