package main

import (
	"awesomego/aftersnow/core"
	"flag"
	"fmt"
)

func main() {
	var Url string
	var filename string
	var ThreadNumber int
	fmt.Println("(-U=<targetUrl> | -F=<target File> | -T=<threads>)")
	flag.StringVar(&filename, "F", "", "filename默认为空")
	flag.IntVar(&ThreadNumber, "T", 10, "线程默认为10")
	flag.StringVar(&Url, "U", "", "Url单独探测默认为空")
	flag.Parse()
	if Url == "" && filename != "" {
		core.Start(filename, ThreadNumber)
		core.Wait()
		core.End()
	} else if Url != "" && filename == "" {
		core.ScanUrl(Url)
	} else {
		fmt.Println("Please input according to the treaty of old King, may the sun guide your direction\n\n")

	}
}
