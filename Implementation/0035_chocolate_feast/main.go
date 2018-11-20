/*
URL			: https://www.hackerrank.com/challenges/chocolate-feast/problem
AUTHOR		: shashank21j
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

// Complete the chocolateFeast function below.
func chocolateFeast(n int32, c int32, m int32) int32 {


}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024 * 1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024 * 1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		ncm := strings.Split(readLine(reader), " ")

		nTemp, err := strconv.ParseInt(ncm[0], 10, 64)
		checkError(err)
		n := int32(nTemp)

		cTemp, err := strconv.ParseInt(ncm[1], 10, 64)
		checkError(err)
		c := int32(cTemp)

		mTemp, err := strconv.ParseInt(ncm[2], 10, 64)
		checkError(err)
		m := int32(mTemp)

		result := chocolateFeast(n, c, m)

		fmt.Fprintf(writer, "%d\n", result)
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
