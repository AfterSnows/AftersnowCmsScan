package utils

import (
	"bufio"
	"fmt"
	"os"
)

type File struct {
	Fp   *os.File
	Scan *bufio.Scanner
}

func Open(filename string) (lines []string) {
	file, err := os.Open(filename)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	return lines
}

func SplitUrls(Line []string, ThreadNumber int, time int, b int, c int, d int, model int) []string {
	Thread := ThreadNumber
	lines := Line
	var line []string
	i := time
	switch model {
	case 0:
		line = lines[i*Thread : Thread+i*Thread]
	case 1:
		if i == c {
			line = lines[i*Thread : i*Thread+d]
		} else {
			line = lines[i*Thread : Thread+i*Thread]

		}
	}
	return line
}

func WriteFile(url string, value string) {
	var filepath string
	filepath = "cms扫描结果" + ".txt"
	file, err := os.OpenFile(filepath, os.O_WRONLY|os.O_APPEND|os.O_CREATE, 0666)
	if err != nil {
		fmt.Println("文件打开失败", err)
	}
	defer file.Close()
	write := bufio.NewWriter(file)
	write.WriteString(url + ":\n" + "product扫描结果:" + value + "\n")
	write.Flush()
}
