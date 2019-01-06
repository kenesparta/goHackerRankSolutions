/*
URL			: https://www.hackerrank.com/challenges/game-of-thrones/problem
AUTHOR		: amititkgp
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

// Complete the gameOfThrones function below.
func gameOfThrones(s string) string {
	var countImp int32
	for _,v := range mapDiffChar(s){
		if v%2 !=0 {
			countImp++
		}
		if countImp > 1 {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	result := gameOfThrones(s)

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
