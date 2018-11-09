package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func filterString(s string) (fs string) {
	var (
		first  string
		second string
	)
	for i, v := range s {
		if i%2 == 0 {
			first += string(v)
		} else {
			second += string(v)
		}
	}
	fs = first + " " + second
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)
	for tItr := 0; tItr < int(t); tItr++ {
		fmt.Printf("%s\n", filterString(readLine(reader)))
	}
}

func readLine(reader *bufio.Reader) string {
	str, _, err := reader.ReadLine()
	if err == io.EOF {
		return ""
	}

	return strings.TrimRight(string(str), "\r\n")
}

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}
