package main

import (
	"fmt"
	"linebot/internal/config"
	_ "linebot/internal/config/db/migrate"
	"linebot/internal/model/product"
	"linebot/internal/route"

	_ "github.com/lib/pq"
)

func main() {
	//model.InitDbContext()
	//db.InitDbContext()
	TestData()
	ginroute := route.InitRouter()
	fmt.Printf("Address: http://localhost:%s/ \n", config.HttpPort)
	ginroute.Run(":" + config.HttpPort)

	//first page
}

func TestData() {
	var p1 = product.Product{
		CampRoundName: "A區",
		Size:          "5m*5m",
		ImageUri:      []string{"https://i.imgur.com/XXwY96T.jpeg", "https://i.imgur.com/3dthZKo.jpeg"},
	}

	var p2 = product.Product{
		CampRoundName: "B區",
		Size:          "5m*5m",
		ImageUri:      []string{"https://i.imgur.com/XXwY96T.jpeg", "https://i.imgur.com/3dthZKo.jpeg"},
	}
	p1.Add()
	p2.Add()
	all, _ := product.GetAll()
	fmt.Println(all)
}
