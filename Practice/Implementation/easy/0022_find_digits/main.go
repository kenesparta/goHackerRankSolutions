package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func breakDigits(n int32) (digits []int) {
	quotient := n
	for quotient > 0 {
		rest := quotient % 10
		quotient = int32(quotient / 10)
		digits = append(digits, int(rest))
	}
	return
}

// Complete the findDigits function below.
func findDigits(n int32) int32 {
	digits := breakDigits(n)
	fmt.Println(digits)
	var countDigits int32
	for _, v := range digits {
		if v != 0 {
			if n%int32(v) == 0 {
				countDigits++
			}
		}

	}
	return countDigits
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

		result := findDigits(n)

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
