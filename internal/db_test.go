package internal

import (
	"fmt"
	"linebot/internal/config"
	_ "linebot/internal/config/db/migrate"
	"testing"
)

func Test_Any(m *testing.T) {
	fmt.Println(config.Config)

	// var x = product.Product{
	// 	CampRoundName: "lalala",
	// }

	// product.AddN(x)

	// r, _ := product.GetAll()
	// fmt.Println(r)
	// product.Add(&x)
	// z, _ := product.GetAll()
	// fmt.Println(z)
}
