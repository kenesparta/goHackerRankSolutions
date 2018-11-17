package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// minNotZero calculates the mininum positive value on array
func minNotZeroAndSubs(arr []int32) (minValueArr []int32, quantity int32) {
	var minValue int32
	// Select non zero value
	for _, v := range arr {
		if v > 0 {
			minValue = v
			break
		}
	}

	for _, v := range arr {
		if v > 0 && minValue >= v {
			minValue = v
		}
	}

	for _, v := range arr {
		rest := v - minValue
		if rest >= 0 {
			quantity++

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
	arrayresult := arr
	for !verifyAllZero(arrayresult) {
		arrayresult, quantity = minNotZeroAndSubs(arrayresult)
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
