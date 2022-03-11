package model

import (
	"regexp"
	"strings"
)

func MatchBody(resp Response, rule FoFaRule, captureType int) bool {
	var re string
	switch captureType {
	case 0:
		re = `(?im)<\s*banner.*>(.*?)<\s*/\s*banner>`
	case 1:
		re = `(?im)<\s*title.*>(.*?)<\s*/\s*title>`
	}
	Results := RegexpBody(resp, re)
	if len(Results) == 0 {
		return false
	}
	for _, Result := range Results {
		if !strings.Contains(
			strings.ToLower(Result),
			strings.ToLower(rule.Content),
		) {
			return false
		}
	}
	return true

}

func RegexpBody(resp Response, re string) []string {
	ret := regexp.MustCompile(re)
	return ret.FindAllString(string(resp.Body), -1)
}

func MatchPort(resp Response, port string) bool {
	/*var h []string
	s := resp.Url
	u, err := url.Parse(s)
	if err != nil {
		return false
	}
	matched, err := regexp.MatchString("((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})(\\.((2(5[0-5]|[0-4]\\d))|[0-1]?\\d{1,2})){3}", u.Host)
	if matched {
		h = strings.Split(u.Host, ":")
		if h[1] == port {
			return true
		}
	}

	*/
	return false
}
