package checkfeatures

import (
	"fmt"
	"regexp"
)

func Check(Match string, Content string, Body string, Server string, Header string) int {
	Result := 1
	switch Match {
	case "banner_contains":
		Result = Result * (Getbanner(Content, Body))
	case "title_contains":
		Result = Result * (Gettitle(Content, Body))
	case "title":
		Result = Result * (Gettitle(Content, Body))
	case "server_contains":
		if Server != "Null" {
			Result = Result * (Getserver(Content, Server))
		}
	case "server":
		if Server != "Null" {
			Result = Result * (Getserver(Content, Server))
		}
	case "cert_contains":
		Result = Result * (Getcert(Content, Body))
	case "protocol_contains":
		Result = Result * (Getprotocol(Content, Body))
	case "protocol":
		Result = Result * (Getprotocol(Content, Body))
	case "header_contains":
		Result = Result * (Getheader(Content, Header))
	case "port_contains":
		Result = Result * (Getport(Content, Body))
	case "body_contains":
		Result = Result * (Getbody(Content, Body))

	default:
		fmt.Println("error", Content)
	}
	return Result

}

func Getbody(Content string, Body string) int {
	matched, _ := regexp.MatchString(Content, Body)
	if matched == true {
		return 1
	}
	return 0
}

func Getbanner(Content string, Body string) int {
	//	ret := regexp.MustCompile(`(?im)<|s*` + Content + `.*>(.*?)<}s*/Is*` + Content + `>`)
	//	result := ret.FindAllStringSubmatch(Body, -1)
	//	for i, subStr := range result {
	//		if Content == subStr[i] {
	//			return 1
	//			fmt.Println(Content)
	//		}
	//	}

	return 0
}

func Gettitle(Content string, Body string) int {

	ret := regexp.MustCompile(`<title>(?s:(.*?))</title>`)
	result := ret.FindAllStringSubmatch(Body, -1)
	for i, subStr := range result {
		if Content == subStr[i] {
			return 1
		}
	}
	return 0
}

func Getcert(Content string, Body string) int {
	return 0
}
func Getprotocol(Content string, Body string) int {
	return 0
}

func Getheader(Content string, Header string) int {
	matched, _ := regexp.MatchString(Content, Header)
	if matched == true {
		return 1
	}
	return 0
}

func Getport(Content string, Body string) int {
	return 0
}

func Getserver(Content string, Server string) int {
	if Content == Server {
		return 1
	}
	return 0
}
