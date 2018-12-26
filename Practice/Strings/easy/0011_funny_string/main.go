/*
URL			: https://www.hackerrank.com/challenges/funny-string/problem
AUTHOR		: Shafaet
DIFFICULTY	: easy
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func abs(n int) int {
	if n < 0 {
		return -1 * n
	} else {
		return n
	}
}

func reverseString(s string) (rev string) {
	for i := len(s) - 1; i >= 0; i-- {
		rev += string(s[i])
	}
	return
}

// Complete the funnyString function below.
func funnyString(s string) string {
	rev := reverseString(s)
	for i := 0; i < len(s)-1; i++ {
		if abs(int(rev[i])-int(rev[i+1])) != abs(int(s[i])-int(s[i+1])) {
			return "Not Funny"
		}
	}
	return "Funny"
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

		result := funnyString(s)

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
