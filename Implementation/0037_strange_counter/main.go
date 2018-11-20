/*
URL			: https://www.hackerrank.com/challenges/strange-code/problem
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

// Complete the strangeCounter function below.
func strangeCounter(t int64) int64 {
	var iBelow int64 = 1
	var iSup int64 = iBelow*2 + 1
	for {
		if t >= iBelow && t <= iSup {
			return (iBelow*2 + 2) - t
		} else {
			iBelow = iBelow*2 + 2
			iSup = iBelow*2 + 1
		}
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	t, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	result := strangeCounter(t)

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
