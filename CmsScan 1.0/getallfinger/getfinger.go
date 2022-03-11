package getallfinger

import (
	"awesomego/aftersnow/checkfeatures"
	"encoding/json"
	"fmt"
	"io/ioutil"
)

type Rule []struct {
	Match   string `json:"match"`
	Content string `json:"content"`
}

type finger struct {
	RuleID         string `json:"rule_id"`
	Level          string `json:"level"`
	Softhard       string `json:"softhard"`
	Product        string `json:"product"`
	Company        string `json:"company"`
	Category       string `json:"category"`
	ParentCategory string `json:"parent_category"`
	Rules          []Rule `json:"rules"`
}

var Ifprint int = 0

func Getfingers(Body string, Server string, Header string) (fingers []finger) {
	filebytes, err := ioutil.ReadFile("aftersnow/fofa.json")
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(filebytes, &fingers)
	for i := 0; i < len(fingers); i++ {
		for j := 0; j < len(fingers[i].Rules); j++ {
			for k := 0; k < len(fingers[i].Rules[j]); k++ {
				Result := checkfeatures.Check(fingers[i].Rules[j][k].Match, fingers[i].Rules[j][k].Content, Body, Server, Header)
				if Result == 1 {
					Ifprint += 1
				} else {
					Ifprint = 0
					break
				}
			}
			if Ifprint == (len(fingers[i].Rules[j])) {
				fmt.Println("Product:", fingers[i].Product)
				break
			} else {
				continue
			}
		}

	}
	return nil
}
