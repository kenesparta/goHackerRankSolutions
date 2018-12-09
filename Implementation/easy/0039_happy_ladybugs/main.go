/*
URL			: https://www.hackerrank.com/challenges/happy-ladybugs/problem
AUTHOR		: nabila_ahmed
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

func hasSpaces(b string) bool {
	for _, v := range b {
		if 95 == v {
			return true
		}
	}
	return false
}

func isLadybugAlone(b string) bool {
	b = strings.Replace(b, "_", "", -1)
	for b != "" {
		cuntTypes := 0
		for _, v2 := range b[1:] {
			if string(b[0]) == string(v2) {
				cuntTypes++
			}
		}
		b = strings.Replace(b, string(b[0]), "", -1)
		// There are a ladybug alone! =(
		if cuntTypes == 0 {
			return true
		}
	}
	return false
}

func hasAdjacent(b string) bool {
	ladybugs := len(b)
	// First ladybug is alone
	if b[1] != b[0] {
		return false
	}

	// Last ladybug is alone
	if b[ladybugs-1] != b[ladybugs-2] {
		return false
	}
	// The others ladybugs
	for i := 1; i < ladybugs-1; i++ {
		if b[i-1] != b[i] && b[i+1] != b[i] {
			return false
		}
	}
	return true
}

// Complete the happyLadybugs function below.
func happyLadybugs(n int32, b string) string {
	if isLadybugAlone(b) {
		return "NO"
	} else if !hasSpaces(b) {
		if hasAdjacent(b) {
			return "YES"
		} else {
			return "NO"
		}
	} else {
		return "YES"
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	gTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	g := int32(gTemp)

	for gItr := 1; gItr <= int(g); gItr++ {
		nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
		checkError(err)

		n := int32(nTemp)

		b := readLine(reader)

		result := happyLadybugs(n, b)

		fmt.Fprintf(writer, "%s\n", result)
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
