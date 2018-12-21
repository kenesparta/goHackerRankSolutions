/*
URL			: https://www.hackerrank.com/challenges/hackerrank-in-a-string/problem
AUTHOR		: Shafaet
DIFFICULTY	: easy
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"regexp"
	"strconv"
	"strings"
)

// Complete the hackerrankInString function below.
func hackerrankInString(s string) string {
	match, _ := regexp.MatchString(`(\w*)h(\w*)c(\w*)k(\w*)e(\w*)r(\w*)r(\w*)a(\w*)n(\w*)k(\w*)`, s)
	if match {
		return "YES"
	} else {
		return "NO"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		result := hackerrankInString(s)

		fmt.Fprintf(writer, "%s\n", result)
	}

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
