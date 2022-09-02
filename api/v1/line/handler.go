package line

import (
	"linebot/internal/model/product"
	"linebot/internal/model/stock"
	"linebot/pkg/tool"
	"log"
)

func SearchRemainCamp(t Search_Time) (r_c []RemainCamp) {
	var err error
	products, err := product.GetAll()
	if err != nil {
		log.Println("Get Products Failed", err)
	}

	for _, p := range products {
		var tmp RemainCamp

		tmp.Product = p
		tmp.Stocks, err = stock.GetStocks_By_ID_and_DateRange(tmp.Product.ID, t.Start, t.End)
		if err != nil {
			log.Println("GetStocks Failed", err)
		}
		var remain []int
		for _, s := range tmp.Stocks {
			remain = append(remain, s.RemainNum)
		}
		//找到最小剩餘數
		tmp.RemainMinAmount, _ = tool.Find_Min_and_Max(remain)

		r_c = append(r_c, tmp)
	}

	return r_c
}

type RemainCamp struct {
	Product         product.Product
	Stocks          []stock.Stock
	RemainMinAmount int
}
