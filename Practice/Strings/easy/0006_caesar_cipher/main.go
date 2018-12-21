/*
URL			: https://www.hackerrank.com/challenges/caesar-cipher-1/problem
AUTHOR		: shashank21j
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

// Uppercase 65 - 90
// lowercase 97 -122

func applyCipher(s rune, k int32) (cipher rune) {
	if 64 <= s && s <= 90 {
		cipher = (int32(s)-65+k)%26 + 65
	} else if 97 <= s && s <= 122 {
		cipher = (int32(s)-97+k)%26 + 97
	} else {
		cipher = s
	}
	return
}

// Complete the caesarCipher function below.
func caesarCipher(n int32, s string, k int32) (newSting string) {
	for _, v := range s {
		newSting += string(applyCipher(v, k))
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

	s := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := caesarCipher(n, s, k)

	fmt.Fprintf(writer, "%s\n", result)

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
