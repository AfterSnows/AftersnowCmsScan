package main

import (
	"awesomego/aftersnow/getallfinger"
	"awesomego/aftersnow/request"
	"bufio"
	"fmt"
	"io"
	"log"
	"os"
)

func main() {
	fi, err := os.Open("aftersnow/text.txt")
	if err != nil {
		log.Fatalln(err)
		return
	}
	defer fi.Close()

	br := bufio.NewReader(fi)
	for {
		s, _, c := br.ReadLine()
		if c == io.EOF {
			break
		}
		str := string(s)
		fmt.Println(str)
		getallfinger.Getfingers(request.Requests(str))

	}

}
