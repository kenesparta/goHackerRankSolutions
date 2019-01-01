/*
URL			: https://www.hackerrank.com/challenges/palindrome-index/problem
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
func isPalindrome(s string) bool {
	comparations := len(s) / 2
	topCom := len(s) - 1
	for i := 0; i < comparations; i++ {
		if s[i] != s[topCom] {
			return false
		} else {
			topCom--
		}
	}
	return true
}

// Complete the palindromeIndex function below.
func palindromeIndex(s string) int32 {
	if isPalindrome(s) {
		return -1
	} else {
		size := int32(len(s))
		topCom := size - 1
		for i := int32(0); i < size/2; i++ {
			palTestBegin := s[0:i] + s[i+1:size]
			if isPalindrome(palTestBegin) {
				return i
			}
			paltestEnd := s[0:topCom] + s[topCom+1:size]
			if isPalindrome(paltestEnd) {
				return topCom
			}
			topCom--
		}
		return -1
	}
	return 0
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

		result := palindromeIndex(s)

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
