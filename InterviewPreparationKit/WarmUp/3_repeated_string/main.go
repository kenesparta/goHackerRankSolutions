package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the repeatedString function below.
func repeatedString(s string, n int64) int64 {
	sizeString := int64(len(s))
	totalParts := int64(n / sizeString)
	restPart := n % sizeString
	var numberOfA int64
	for _, v := range s {
		if v == 'a' {
			numberOfA++
		}
	}

	var restOfA int64
	for i := 0; int64(i) < restPart; i++ {
		if s[i] == 'a' {
			restOfA++
		}
	}
	return totalParts*numberOfA + restOfA
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	n, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := repeatedString(s, n)

	fmt.Fprintf(writer, "%d\n", result)

	writer.Flush()
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
