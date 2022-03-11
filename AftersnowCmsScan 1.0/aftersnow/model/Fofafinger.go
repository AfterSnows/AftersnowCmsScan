package model

import (
	"awesomego/aftersnow/assets"
	"encoding/json"
	"strings"
)

type FoFaRule struct {
	Match   string `json:"match"`
	Content string `json:"content"`
}

type FoFaFingers struct {
	RuleID         string       `json:"rule_id"`
	Level          string       `json:"level"`
	Softhard       string       `json:"softhard"`
	Product        string       `json:"product"`
	Company        string       `json:"company"`
	Category       string       `json:"category"`
	ParentCategory string       `json:"parent_category"`
	Rules          [][]FoFaRule `json:"rules"`
}

type FoFaFinger []FoFaFingers

var FoFa FoFaFinger

func init() {
	err := json.Unmarshal(assets.FingerData, &FoFa) //结构体化
	if err != nil {
		panic(err.Error())
	}
}

func Check(resp Response, fingers FoFaFinger) ([]FoFaFingers, bool) {
	matchFinger := []FoFaFingers{}
	isSubRuleMatched := true
	isAllSubRulesMatched := true
	for _, Finger := range fingers {
		for _, Rule := range Finger.Rules {
			for _, subRule := range Rule {
				a := strings.Split(subRule.Match, "_")[0]
				switch a {
				case "banner":
					isSubRuleMatched = MatchBody(resp, subRule, 0)
				case "title":
					isSubRuleMatched = MatchBody(resp, subRule, 1)
				case "body":
					if !strings.Contains(
						strings.ToLower(string(resp.Body)),
						strings.ToLower(subRule.Content),
					) {
						isSubRuleMatched = false
					}
				case "header":
					if resp.Headers.Get(subRule.Content) == "" {
						isSubRuleMatched = false
					}
				case "server":
					if !strings.Contains(
						strings.ToLower(resp.Headers.Get("Server")),
						strings.ToLower(subRule.Content),
					) {
						isSubRuleMatched = false
					}
				case "cert":
					if (resp.Cert == nil) || (resp.Cert != nil && !strings.Contains(string(resp.Cert), subRule.Content)) {
						isSubRuleMatched = false
					}
				case "port":
					isSubRuleMatched = MatchPort(resp, subRule.Content)
				case "protocol":
					isSubRuleMatched = false
				default:
					isSubRuleMatched = false
				}
				isAllSubRulesMatched = isAllSubRulesMatched && isSubRuleMatched
				if !isAllSubRulesMatched {
					break
				}
			}
			if isAllSubRulesMatched {
				matchFinger = append(matchFinger, Finger)
			}
			isSubRuleMatched = true
			isAllSubRulesMatched = true
		}
	}
	if len(matchFinger) != 0 {
		return matchFinger, true
	}
	return nil, false
}
