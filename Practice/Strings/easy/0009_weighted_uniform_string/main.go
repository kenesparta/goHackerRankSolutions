/*
URL			: https://www.hackerrank.com/challenges/weighted-uniform-string/problem
AUTHOR		: nabila_ahmed
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

// lowercase 97 -122
const rest = int(96)

// Complete the weightedUniformStrings function below.
func createUniformSubString(s string) (wStr []int) {
	size := len(s)
	totalAccSum := 1
	for i := 0; i < size-1; i++ {
		if s[i] == s[i+1] {
			totalAccSum++
		} else {
			totalAccSum = 1
		}
		value := int(s[i]) - rest
		wStr = append(wStr, value*totalAccSum)
	}
	value := int(s[size-1]) - rest
	wStr = append(wStr, value)
	return
}

func weightedUniformStrings(s string, queries []int32) (resp []string) {
	unif := createUniformSubString(s)
	fmt.Println(unif)
	for _, v1 := range queries {
		var flag bool
		for _, v2 := range unif {
			if v1 == int32(v2) {
				flag = true
				break
			}
		}
		if flag {
			resp = append(resp, "Yes")
		} else {
			resp = append(resp, "No")
		}
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	queriesCount, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)

	var queries []int32

	for i := 0; i < int(queriesCount); i++ {
		queriesItemTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		queriesItem := int32(queriesItemTemp)
		queries = append(queries, queriesItem)
	}

	result := weightedUniformStrings(s, queries)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

		if i != len(result)-1 {
			fmt.Fprintf(writer, "\n")
		}
	}

	fmt.Fprintf(writer, "\n")

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
