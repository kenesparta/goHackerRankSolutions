/*
URL			: https://www.hackerrank.com/challenges/cavity-map/problem
AUTHOR		: Gera1d
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

func getNSWE(g []string, i, j int) (n, s, e, w string) {
	return string(g[i-1][j]), string(g[i+1][j]), string(g[i][j+1]), string(g[i][j-1])
}

func hasArround(grid []string, i, j int, value string) bool {
	n, s, e, w := getNSWE(grid, i, j)
	return value == n || value == s || value == e || value == w
}

func isMaxValue(grid []string, i, j int, value string) bool {
	n, s, e, w := getNSWE(grid, i, j)
	return value > n && value > s && value > e && value > w
}

// Complete the cavityMap function below.
func cavityMap(grid []string) []string {
	size := len(grid)
	for i, v := range grid {
		if i != 0 && i != size-1 {
			var newString string
			for j, s := range v {
				if j != 0 && j != size-1 {
					if isMaxValue(grid, i, j, string(s)) && !hasArround(grid, i, j, string(s)) {
						newString += "X"
					} else {
						newString += string(s)
					}
				} else {
					newString += string(s)
				}
			}
			grid[i] = newString
		}
	}
	return grid
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

	var grid []string

	for i := 0; i < int(n); i++ {
		gridItem := readLine(reader)
		grid = append(grid, gridItem)
	}

	result := cavityMap(grid)

	for i, resultItem := range result {
		fmt.Fprintf(writer, "%s", resultItem)

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
