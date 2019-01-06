/*
URL			: https://www.hackerrank.com/challenges/anagram/problem
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

// Mapping different strings
func mapDiffChar(s string) map[string]int32 {
	counts := make(map[string]int32)
	for _, l := range s {
		strL := string(l)
		if _, ok := counts[strL]; ok {
			counts[strL] ++
		} else {
			counts[strL] = 1
		}
	}
	return counts
}

// Count the number of letters on string
func countChars(c, s string) (total int32) {
	for _, v := range s {
		if c == string(v) {
			total++
		}
	}
	return
}

// Complete the anagram function below.
func anagram(s string) int32 {
	size := len(s)
	if size%2 != 0 {
		return -1
	} else {
		firstPart := s[0 : size/2]
		secondPart := s[size/2:]
		var totalChanges int32
		for k, v := range mapDiffChar(firstPart) {
			diff := v - countChars(k, secondPart)
			if diff > 0 {
				totalChanges += diff
			}
		}
		return totalChanges
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

		result := anagram(s)

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
