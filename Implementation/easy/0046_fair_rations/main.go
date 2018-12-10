/*
URL			: https://www.hackerrank.com/challenges/fair-rations/problem
AUTHOR		: kevinsogo
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

type Node struct {
	data       int32
	prev, next *Node
}

// Complete the fairRations function below.
func fairRations(B []int32) (loaves int32) {
	var (
		size        = len(B)
		oddQnt      int32
		oddDistance int32
		oddFlag     = false
	)

	// Quantity ODD numbers on list
	for i := 0; i < size; i++ {
		if B[i]%2 != 0 {
			oddQnt++
		}
	}

	// Count loaves
	if oddQnt%2 == 0 {
		for i := 0; i < size; i++ {
			if B[i]%2 != 0 {
				oddQnt++
				oddFlag = true
			}

			// When find other odd element
			if oddQnt%2 == 0 {
				oddFlag = false
				loaves += oddDistance * 2
			}

			// Start sum
			if oddFlag {
				oddDistance++
			} else {
				oddDistance = 0
			}


		}
	} else {
		loaves = -1
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	NTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	N := int32(NTemp)

	BTemp := strings.Split(readLine(reader), " ")

	var B []int32

	for i := 0; i < int(N); i++ {
		BItemTemp, err := strconv.ParseInt(BTemp[i], 10, 64)
		checkError(err)
		BItem := int32(BItemTemp)
		B = append(B, BItem)
	}

	result := fairRations(B)
	if result == -1 {
		fmt.Fprintf(writer, "%s\n", "NO")
	} else {
		fmt.Fprintf(writer, "%d\n", result)
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
