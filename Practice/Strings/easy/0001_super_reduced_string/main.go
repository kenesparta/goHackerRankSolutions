/*
URL			: https://www.hackerrank.com/challenges/reduced-string/problem
AUTHOR		: harshil7924
DIFFICULTY	: easy
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

// Complete the superReducedString function below.
func superReducedString(s string) string {
	var (
		hIdx = 1 // Head Index
		tIdx int // Tail index
		size = len(s)
	)
	for hIdx < size {
		if s[tIdx] == s[hIdx] {
			s = s[0:tIdx] + s[hIdx+1:]
			hIdx = 1
			tIdx = 0
		} else {
			hIdx++
			tIdx++
		}
		size = len(s)
	}
	if s == "" {
		return "Empty String"
	} else {
		return s
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	s := readLine(reader)

	result := superReducedString(s)

	fmt.Fprintf(writer, "%s\n", result)

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
