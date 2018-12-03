package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func isPrime(number int64) bool {
	if number == 1 {
		return false
	}

	if number%2 == 0 && number > 2 {
		return false
	}

	var i int64 = 3
	for i*i <= number {

		if number%i == 0 {
			return false
		}
		i += 2
	}
	return true
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int(nTemp)

	for i := 0; i < n; i++ {
		item, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		if isPrime(item) {
			fmt.Println("Prime")
		} else {
			fmt.Println("Not prime")
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
