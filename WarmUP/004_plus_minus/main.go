package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the plusMinus function below.
func plusMinus(arr []int32) {
	var positives int32
	var negatives int32
	var zeros int32
	total := len(arr)
	for i := 0; i < total; i++ {
		if arr[i] > 0 {
			positives++
		} else if arr[i] < 0 {
			negatives++
		} else {
			zeros++
		}
	}
	fP := float32(positives) / float32(total)
	nP := float32(negatives) / float32(total)
	zP := float32(zeros) / float32(total)
	fmt.Printf("%.6f\n", fP)
	fmt.Printf("%.6f\n", nP)
	fmt.Printf("%.6f\n", zP)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

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

	plusMinus(arr)
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
