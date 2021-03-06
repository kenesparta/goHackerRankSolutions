/*
URL			: https://www.hackerrank.com/challenges/cut-the-sticks/problem
AUTHOR		: shashank21j
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

// minNotZero calculates the minimum positive value on array
func minNotZeroAndSubs(arr []int32) (minValueArr []int32, changes int32) {
	var minValue int32

	// Select non zero value
	for _, v := range arr {
		if v > 0 {
			minValue = v
			break
		}
	}

	// Select min value
	for _, v := range arr {
		if v > 0 && minValue >= v {
			minValue = v
		}
	}

	// Count the changes
	for _, v := range arr {
		rest := v - minValue
		if rest >= 0 {
			changes++
		} else {
			rest = 0
		}
		minValueArr = append(minValueArr, rest)
	}
	return
}

// verify if all is zero
func verifyAllZero(arr []int32) bool {
	for _, v := range arr {
		if v != 0 {
			return false
		}
	}
	return true
}

// Complete the cutTheSticks function below.
func cutTheSticks(arr []int32) (numberCuts []int32) {
	var quantity int32
	arrayResult := arr
	for !verifyAllZero(arrayResult) {
		arrayResult, quantity = minNotZeroAndSubs(arrayResult)
		numberCuts = append(numberCuts, quantity)
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

	arrTemp := strings.Split(readLine(reader), " ")

	var arr []int32

	for i := 0; i < int(n); i++ {
		arrItemTemp, err := strconv.ParseInt(arrTemp[i], 10, 64)
		checkError(err)
		arrItem := int32(arrItemTemp)
		arr = append(arr, arrItem)
	}

	result := cutTheSticks(arr)

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
