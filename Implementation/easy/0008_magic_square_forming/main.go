package main

import (
	"bufio"
	"fmt"
	"io"
	"math"
	"os"
	"strconv"
	"strings"
)

var magicSquares = [][]int{
	{8, 1, 6, 3, 5, 7, 4, 9, 2},
	{6, 1, 8, 7, 5, 3, 2, 9, 4},
	{4, 9, 2, 3, 5, 7, 8, 1, 6},
	{2, 9, 4, 7, 5, 3, 6, 1, 8},
	{8, 3, 4, 1, 5, 9, 6, 7, 2},
	{4, 3, 8, 9, 5, 1, 2, 7, 6},
	{6, 7, 2, 1, 5, 9, 8, 3, 4},
	{2, 7, 6, 9, 5, 1, 4, 3, 8},
}

func toLinearArray(s [][]int32) (linearS []int32) {
	for _, v1 := range s {
		for _, v2 := range v1 {
			linearS = append(linearS, v2)
		}
	}
	return
}

// Complete the formingMagicSquare function below.
func formingMagicSquare(s [][]int32) (minWeight int32) {
	linearArray := toLinearArray(s)
	for j, m := range magicSquares {
		var weight int32
		for i, v2 := range m {
			weight += int32(math.Abs(float64(v2) - float64(linearArray[i])))
		}
		fmt.Printf(">>%d\n", weight)
		// Start minWeight
		if j == 0 {
			minWeight = weight
		}
		if weight < minWeight {
			minWeight = weight
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

	var s [][]int32
	for i := 0; i < 3; i++ {
		sRowTemp := strings.Split(readLine(reader), " ")

		var sRow []int32
		for _, sRowItem := range sRowTemp {
			sItemTemp, err := strconv.ParseInt(sRowItem, 10, 64)
			checkError(err)
			sItem := int32(sItemTemp)
			sRow = append(sRow, sItem)
		}

		if len(sRow) != 3 {
			panic("Bad input")
		}

		s = append(s, sRow)
	}

	result := formingMagicSquare(s)

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
