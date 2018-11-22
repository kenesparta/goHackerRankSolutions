/*
URL			: https://www.hackerrank.com/challenges/service-lane/problem
AUTHOR		: abhiranjan
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

// Complete the serviceLane function below.
func serviceLane(n int32, cases [][]int32) []int32 {

}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nt := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nt[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	tTemp, err := strconv.ParseInt(nt[1], 10, 64)
	checkError(err)
	t := int32(tTemp)

	widthTemp := strings.Split(readLine(reader), " ")

	var width []int32

	for i := 0; i < int(n); i++ {
		widthItemTemp, err := strconv.ParseInt(widthTemp[i], 10, 64)
		checkError(err)
		widthItem := int32(widthItemTemp)
		width = append(width, widthItem)
	}

	var cases [][]int32
	for i := 0; i < int(t); i++ {
		casesRowTemp := strings.Split(readLine(reader), " ")

		var casesRow []int32
		for _, casesRowItem := range casesRowTemp {
			casesItemTemp, err := strconv.ParseInt(casesRowItem, 10, 64)
			checkError(err)
			casesItem := int32(casesItemTemp)
			casesRow = append(casesRow, casesItem)
		}

		if len(casesRow) != int(2) {
			panic("Bad input")
		}

		cases = append(cases, casesRow)
	}

	result := serviceLane(n, cases)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%d", resultItem)

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
