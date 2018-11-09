package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func convertOneDim(arr [][]int32) (arrOneDim [36]int32) {
	var totalIndex int32
	for _, v1 := range arr {
		for _, v2 := range v1 {
			arrOneDim[totalIndex] = v2
			totalIndex++
		}
	}
	return
}

func sum(arr []int32) (totalSum int32) {
	for _, v := range arr {
		totalSum += v
	}
	return
}

// Complete the hourglassSum function below.
func hourglassSum(arr [][]int32) (maxValue int32) {
	magicNumbers := []int32{0, 1, 2, 3, 6, 7, 8, 9, 12, 13, 14, 15, 18, 19, 20, 21}
	arrOneDim := convertOneDim(arr)
	maxValue = -63
	for _, v := range magicNumbers {
		var sumValue int32
		sumValue += sum(arrOneDim[v:v+3]) + arrOneDim[v+7] + sum(arrOneDim[v+12:v+15])
		if maxValue < sumValue {
			maxValue = sumValue
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

	var arr [][]int32
	for i := 0; i < 6; i++ {
		arrRowTemp := strings.Split(readLine(reader), " ")

		var arrRow []int32
		for _, arrRowItem := range arrRowTemp {
			arrItemTemp, err := strconv.ParseInt(arrRowItem, 10, 64)
			checkError(err)
			arrItem := int32(arrItemTemp)
			arrRow = append(arrRow, arrItem)
		}

		if len(arrRow) != int(6) {
			panic("Bad input")
		}

		arr = append(arr, arrRow)
	}

	result := hourglassSum(arr)

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
