/*
URL			: https://www.hackerrank.com/challenges/minimum-distances/problem
AUTHOR		: Shafaet
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

// Complete the minimumDistances function below.
func minimumDistances(a []int32) (minDist int32) {
	var countEquals int32
	size := len(a)
	minDist = int32(size)
	for i := 0; i < size-1; i++ {
		for j := i + 1; j < size; j++ {
			if a[i] == a[j] {
				// Calculate the distance
				d := int32(j - i)
				if d < minDist {
					minDist = d
				}
				countEquals++
			}
		}
	}
	if countEquals == 0 {
		minDist = -1
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

	aTemp := strings.Split(readLine(reader), " ")

	var a []int32

	for i := 0; i < int(n); i++ {
		aItemTemp, err := strconv.ParseInt(aTemp[i], 10, 64)
		checkError(err)
		aItem := int32(aItemTemp)
		a = append(a, aItem)
	}

	result := minimumDistances(a)

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
