/*
URL			: https://www.hackerrank.com/challenges/gem-stones/problem
AUTHOR		: darkshadows
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

// Verify if an current letter is in array
func inArray(s string, letter string) bool {
	for _, v := range s {
		if string(v) == letter {
			return true
		}
	}
	return false
}

// Clean all the letters that repeat on a given string
func cleanRep(s string) string {
	banList := string(s[0])
	for _, v := range s {
		if !inArray(banList, string(v)) {
			banList += string(v)
		}
	}
	return banList
}

// Complete the gemstones function below.
func gemstones(arr []string) (gems int32) {
	pivotalArray := cleanRep(arr[0])
	//fmt.Println(pivotalArray)
	for _, v1 := range pivotalArray {
		flag := true
		for _, v2 := range arr[1:] {
			if !inArray(v2, string(v1)) {
				flag = false
				break
			}
		}
		//fmt.Println(flag)
		if flag {
			gems++
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

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)

	var arr []string

	for i := 0; i < int(n); i++ {
		arrItem := readLine(reader)
		arr = append(arr, arrItem)
	}

	result := gemstones(arr)

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
