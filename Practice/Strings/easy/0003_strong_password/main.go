/*
URL			: https://www.hackerrank.com/challenges/strong-password/problem
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

const specialCharacters = "!@#$%^&*()-+"

func hasElement(password string, inf, sup rune) int32 {
	for _, v := range password {
		if v >= inf && v <= sup {
			return 0
		}
	}
	return 1
}

func hasSpecialChar(password string) int32 {
	for _, v := range password {
		for _, s := range specialCharacters {
			if s == v {
				return 0
			}
		}
	}
	return 1
}

// Complete the minimumNumber function below.
func minimumNumber(n int32, password string) int32 {
	// Return the minimum number of characters to make the password strong
	minLength := 6 - n
	nNumbers := hasElement(password, 47, 57)
	nUppChar := hasElement(password, 65, 90)
	nLowChar := hasElement(password, 97, 122)
	nSpecialChar := hasSpecialChar(password)
	minReq := nNumbers + nUppChar + nLowChar + nSpecialChar
	if minLength > minReq {
		return minLength
	} else {
		return minReq
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
