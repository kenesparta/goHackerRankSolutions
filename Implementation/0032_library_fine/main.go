/*
URL			: https://www.hackerrank.com/challenges/library-fine/problem
AUTHOR		: vatsalchanana
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

// Complete the libraryFine function below.
func libraryFine(d1 int32, m1 int32, y1 int32, d2 int32, m2 int32, y2 int32) (int32) {
	dayDiff := d1 - d2
	monthDiff := m1 - m2
	yearDiff := y1 - y2
	if yearDiff < 0 {
		return 0
	} else if yearDiff == 0 {
		if monthDiff < 0 {
			return 0
		} else if monthDiff == 0 {
			if dayDiff <= 0 {
				return 0
			} else {
				return dayDiff * 15
			}
		} else {
			return monthDiff * 500
		}
	} else {
		return 10000
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	d1M1Y1 := strings.Split(readLine(reader), " ")

	d1Temp, err := strconv.ParseInt(d1M1Y1[0], 10, 64)
	checkError(err)
	d1 := int32(d1Temp)

	m1Temp, err := strconv.ParseInt(d1M1Y1[1], 10, 64)
	checkError(err)
	m1 := int32(m1Temp)

	y1Temp, err := strconv.ParseInt(d1M1Y1[2], 10, 64)
	checkError(err)
	y1 := int32(y1Temp)

	d2M2Y2 := strings.Split(readLine(reader), " ")

	d2Temp, err := strconv.ParseInt(d2M2Y2[0], 10, 64)
	checkError(err)
	d2 := int32(d2Temp)

	m2Temp, err := strconv.ParseInt(d2M2Y2[1], 10, 64)
	checkError(err)
	m2 := int32(m2Temp)

	y2Temp, err := strconv.ParseInt(d2M2Y2[2], 10, 64)
	checkError(err)
	y2 := int32(y2Temp)

	result := libraryFine(d1, m1, y1, d2, m2, y2)

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
