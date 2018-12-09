/*
URL			: https://www.hackerrank.com/challenges/equality-in-a-array/problem
AUTHOR		: muratekici
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

// Complete the equalizeArray function below.
func equalizeArray(arr []int32) (minEqualizer int32) {
	// Count the repetitions of each element within the array
	sizeArray := int32(len(arr))
	minEqualizer = sizeArray
	for len(arr) > 1 {
		var (
			newArr      []int32
			count       int32
			countEquals int32
		)
		firstElem := arr[0]

		// Compare the numbers
		for _, v2 := range arr[1:] {
			if firstElem == v2 {
				countEquals++
			} else {
				newArr = append(newArr, v2)
			}
		}
		qntRest := sizeArray - (countEquals + 1)
		if minEqualizer > qntRest {
			minEqualizer = qntRest
		}
		arr = newArr
		count++
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

	result := equalizeArray(arr)

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
