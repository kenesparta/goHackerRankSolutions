/*
URL			: https://www.hackerrank.com/challenges/manasa-and-stones/problem
AUTHOR		: amititkgp
DIFFICULTY	: easy
*/

package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
	"strconv"
	"strings"
)

// Complete the stones function below.
func stones(n int32, a int32, b int32) (stoneValue []int32) {
	var (
		init = n - 1
		add  int32
	)
	// Set the first value
	stoneValue = append(stoneValue, init*a)
	add = 1
	for i := init - 1; i >= 0; i-- {
		result := i*a + b*add
		size := len(stoneValue)
		if result != stoneValue[size-1] {
			stoneValue = append(stoneValue, result)
		}
		add++
	}
	sort.Slice(stoneValue, func(i, j int) bool {
		return stoneValue[i] < stoneValue[j]
	})
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	TTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	T := int32(TTemp)

	for TItr := 0; TItr < int(T); TItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		n := int32(nTemp)

		aTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		a := int32(aTemp)

		bTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)
		b := int32(bTemp)

		result := stones(n, a, b)

		for i, resultItem := range result {
			fmt.Fprintf(writer, "%d", resultItem)

			if i != len(result)-1 {
				fmt.Fprintf(writer, " ")
			}
		}

		fmt.Fprintf(writer, "\n")
	}

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
