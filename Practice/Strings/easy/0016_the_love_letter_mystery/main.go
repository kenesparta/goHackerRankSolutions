/*
URL			: https://www.hackerrank.com/challenges/the-love-letter-mystery/problem
AUTHOR		: amititkgp
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

// Palindrome better solution
func abs(n int32) int32 {
	if n < 0 {
		return n * -1
	} else {
		return n
	}
}

// Complete the theLoveLetterMystery function below.
func theLoveLetterMystery(s string) int32 {
	var (
		nComp  = len(s) / 2
		topCom = len(s) - 1
		minOp  int32
	)
	for i := 0; i < nComp; i++ {
		dif := abs(int32(s[i]) - int32(s[topCom]))
		if dif != 0 {
			minOp += dif
		}
		topCom--
	}
	return minOp
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

		result := theLoveLetterMystery(s)

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
