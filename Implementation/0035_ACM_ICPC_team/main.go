/*
URL			: https://www.hackerrank.com/challenges/acm-icpc-team/problem
AUTHOR		: dheeraj
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

// Count the max value occurences on array
func countMax(arr []int32) []int32 {
	var countMax int32

	// Get the max value
	maxVal := arr[0]
	for _, v := range arr {
		if maxVal <= v {
			maxVal = v
		}
	}

	// How many "max values" are in the array?
	for _, v := range arr {
		if v == maxVal {
			countMax++
		}
	}
	return []int32{maxVal, countMax}
}

// Complete the acmTeam function below.
func acmTeam(topic []string) []int32 {
	sizeSubjects := len(topic[0])
	var know []int32
	for i, v1 := range topic {
		for _, v2 := range topic[i+1:] {
			var countKnows int32
			for k := 0; k < sizeSubjects; k++ {
				if string(v1[k]) == "1" || string(v2[k]) == "1" {
					countKnows++
				}
			}
			know = append(know, countKnows)
		}
	}
	return countMax(know)
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nm := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nm[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	// mTemp, err := strconv.ParseInt(nm[1], 10, 64)
	// checkError(err)
	// m := int32(mTemp)

	var topic []string

	for i := 0; i < int(n); i++ {
		topicItem := readLine(reader)
		topic = append(topic, topicItem)
	}

	result := acmTeam(topic)

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
