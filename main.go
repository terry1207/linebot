package main

import (
	"fmt"
	"linebot/config"

	"linebot/internal/route"

	"linebot/internal/repository"

	_ "github.com/lib/pq"
)

func main() {
	repository.InitDbContext()

	ginroute := route.InitRouter()

	ginroute.Run(":" + config.HttpPort)
	fmt.Println("Hi")
}
