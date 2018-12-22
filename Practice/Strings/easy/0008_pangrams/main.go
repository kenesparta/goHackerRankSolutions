/*
URL			: https://www.hackerrank.com/challenges/pangrams/problem
AUTHOR		: Bidhan
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

func inArray(s []string, letter string) bool {
	for _, v := range s {
		if v == letter {
			return true
		}
	}
	return false
}

func cleanAndCountRep(s string) int32 {
	banList := []string{string(s[0])}
	var count int32
	for _, v := range s {
		if !inArray(banList, string(v)) {
			count++
			banList = append(banList, string(v))
		}
	}
	return count + 1
}

// Complete the pangrams function below.
func pangrams(s string) string {
	if cleanAndCountRep(strings.ToLower(s)) == 27 {
		return "pangram"
	} else {
		return "not pangram"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := pangrams(s)

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
