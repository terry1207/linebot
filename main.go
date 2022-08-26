package main

import (
	"fmt"
	"linebot/internal/config"
	_ "linebot/internal/config/db/Migrations"
	"linebot/internal/route"

	_ "github.com/lib/pq"
)

func main() {
	//model.InitDbContext()
	//db.InitDbContext()

	ginroute := route.InitRouter()
	fmt.Printf("Address: http://localhost:%s/ \n", config.HttpPort)
	ginroute.Run(":" + config.HttpPort)

	//first page
}
