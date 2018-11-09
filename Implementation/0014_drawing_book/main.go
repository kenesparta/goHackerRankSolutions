package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the pageCount function below.
 */
func pageCount(n int32, p int32) (counts int32) {
	var fromEnd int32
	var fromStart int32
	if n%2 != 0 {
		fromEnd = int32((n - p) / 2)
		fromStart = int32(p / 2)
	} else {
		fromEnd = int32((n + 1 - p) / 2)
		fromStart = int32(p / 2)
	}

	if fromEnd < fromStart {
		counts = fromEnd
	} else {
		counts = fromStart
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	pTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	p := int32(pTemp)

	result := pageCount(n, p)

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
