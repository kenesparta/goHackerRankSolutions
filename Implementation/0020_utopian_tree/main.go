package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

// Complete the utopianTree function below.
func utopianTree(n int32) (grow int32) {
	if n%2 == 0 {
		grow = int32(math.Pow(2, float64(n/2+1)) - 1)
	} else {
		grow = 2 * int32(math.Pow(2, float64((n-1)/2+1))-1)
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int32(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		result := utopianTree(n)

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
