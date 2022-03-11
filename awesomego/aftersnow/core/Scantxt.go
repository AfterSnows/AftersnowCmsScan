package core

import (
	"awesomego/aftersnow/model"
	"awesomego/aftersnow/utils"
	"github.com/panjf2000/ants"
	"log"
	"strings"
	"sync"
)

var Fetch *ants.PoolWithFunc
var WG sync.WaitGroup

func Task(targetsMap interface{}) {
	defer WG.Done()
	targets, ok := targetsMap.([]string)
	if !ok {
		return
	}

	for _, target := range targets {
		resp := model.Request(target)
		datas, isFound := model.Check(*resp, model.FoFa)
		if isFound {
			products := make([]string, 0)
			for _, data := range datas {
				products = append(products, data.Product)
			}
			url := resp.Url
			Product := strings.Join(products, ",")
			log.Printf("URL:%s => %v\n", url, Product)
			utils.WriteFile(url, Product)
		} else {
			log.Printf("URL:%s => not find\n", resp.Url)
		}
	}

}

func Start(filename string, ThreadNumber int) {
	var err error

	Fetch, err = ants.NewPoolWithFunc(ThreadNumber, Task)
	if err != nil {
		log.Fatal(err.Error())
	}
	a := utils.Open(filename)
	UrlLine := a
	b := len(a)
	c := b / ThreadNumber //循环次数
	d := b % ThreadNumber //最后一次次数数组
	if d == 0 {
		for i := 0; i < c; i++ {
			ThreadTarget := utils.SplitUrls(UrlLine, ThreadNumber, i, b, c, d, 0)
			WG.Add(1)
			if err := Fetch.Invoke(ThreadTarget); err != nil {
				log.Fatal(err.Error())
			}
		}
	} else {
		for i := 0; i <= c; i++ {
			ThreadTarget := utils.SplitUrls(UrlLine, ThreadNumber, i, b, c, d, 1)
			WG.Add(1)
			if err := Fetch.Invoke(ThreadTarget); err != nil {
				log.Fatal(err.Error())
			}
		}
	}

}

func Wait() {
	WG.Wait()
}

func End() {
	Fetch.Release()
}
