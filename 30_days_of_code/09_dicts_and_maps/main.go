package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 2048*2048)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := nTemp

	arr := make(map[string]string)

	for i := int64(0); i < n; i++ {
		arrTemp := strings.Split(readLine(reader), " ")
		arr[arrTemp[0]] = arrTemp[1]
	}

	for i := int64(0); i < n; i++ {
		read := readLine(reader)
		if read == "" {
			break
		}
		if value, w := arr[read]; w {
			fmt.Printf("%s=%s\n", read, value)
		} else {
			fmt.Println("Not found")
		}
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
