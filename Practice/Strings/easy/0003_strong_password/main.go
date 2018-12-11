/*
URL			: https://www.hackerrank.com/challenges/strong-password/problem
AUTHOR		: anveshi
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

const (
	numbers           = "0123456789"
	lowerCase         = "abcdefghijklmnopqrstuvwxyz"
	upperCase         = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	specialCharacters = "!@#$%^&*()-+"
)

// Complete the minimumNumber function below.
func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	var minStrong int32
	if n <= 6 {
		minStrong = n-6
	}

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

	password := readLine(reader)

	answer := minimumNumber(n, password)

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
