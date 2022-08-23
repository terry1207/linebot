package main

import (
	"linebot/config"

	"linebot/internal/repository/db"
	"linebot/internal/route"

	_ "github.com/lib/pq"
)

func main() {
	//repository.InitDbContext()
	db.InitDbContext()

	ginroute := route.InitRouter()

	ginroute.Run(":" + config.HttpPort)

	//first page
}
