package core

import (
	"awesomego/aftersnow/model"
	"log"
	"strings"
)

func ScanUrl(url string) {
	resp := model.Request(url)
	datas, isFound := model.Check(*resp, model.FoFa)
	if isFound {
		products := make([]string, 0)
		for _, data := range datas {
			products = append(products, data.Product)
		}

		log.Printf("URL:%s => %v\n", resp.Url, strings.Join(products, ","))
	} else {
		log.Printf("URL:%s => not find\n", resp.Url)
	}
}
