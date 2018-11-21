package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sumSeparateDigits(n int64) int64 {
	s := strconv.FormatInt(n, 10)
	midString := len(s) / 2
	left, _ := strconv.ParseInt(string(s[0:midString]), 10, 64)
	right, _ := strconv.ParseInt(string(s[midString:]), 10, 64)
	return int64(right + left)

}

// Complete the kaprekarNumbers function below.
func kaprekarNumbers(p, q int64) {
	var count int64
	for i := p; i <= q; i++ {
		if i == sumSeparateDigits(i*i) {
			count++
			fmt.Printf("%d ", i)
		}
	}
	if count == 0 {
		fmt.Print("INVALID RANGE")
	}
	fmt.Println()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int64(pTemp)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int64(qTemp)

	kaprekarNumbers(p, q)
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
