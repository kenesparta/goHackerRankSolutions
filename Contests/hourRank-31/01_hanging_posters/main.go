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

/*
 * Complete the 'solve' function below.
 *
 * The function is expected to return an INTEGER.
 * The function accepts following parameters:
 *  1. INTEGER h
 *  2. INTEGER_ARRAY wallPoints
 *  3. INTEGER_ARRAY lengths
 */

func solve(h int32, wallPoints []int32, lengths []int32) int32 {
	// Write your code here
	sizeArr := len(wallPoints)
	var maxVal float64
	for i := 0; i < sizeArr; i++ {
		heightAchieve := float64(wallPoints[i]-h) - .25*float64(lengths[i])
		heightAchieve = math.Ceil(heightAchieve)
		if maxVal < heightAchieve {
			maxVal = heightAchieve
		}
	}
	return int32(maxVal)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	firstMultipleInput := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	nTemp, err := strconv.ParseInt(firstMultipleInput[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	hTemp, err := strconv.ParseInt(firstMultipleInput[1], 10, 64)
	checkError(err)
	h := int32(hTemp)

	wallPointsTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var wallPoints []int32

	for i := 0; i < int(n); i++ {
		wallPointsItemTemp, err := strconv.ParseInt(wallPointsTemp[i], 10, 64)
		checkError(err)
		wallPointsItem := int32(wallPointsItemTemp)
		wallPoints = append(wallPoints, wallPointsItem)
	}

	lengthsTemp := strings.Split(strings.TrimSpace(readLine(reader)), " ")

	var lengths []int32

	for i := 0; i < int(n); i++ {
		lengthsItemTemp, err := strconv.ParseInt(lengthsTemp[i], 10, 64)
		checkError(err)
		lengthsItem := int32(lengthsItemTemp)
		lengths = append(lengths, lengthsItem)
	}

	answer := solve(h, wallPoints, lengths)

	fmt.Fprintf(writer, "%d\n", answer)

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
