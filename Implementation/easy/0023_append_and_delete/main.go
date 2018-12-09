package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func min(a int32, b int32) (minNumber int32) {
	if a < b {
		minNumber = a
	} else {
		minNumber = b
	}
	return
}

// Complete the appendAndDelete function below.
func appendAndDelete(s string, t string, k int32) (response string) {
	var (
		sizeS     = int32(len(s))
		sizeT     = int32(len(t))
		countSame int32
	)
	for i := int32(0); i < min(sizeS, sizeT); i++ {
		if string(s[i]) == string(t[i]) {
			countSame++
		} else {
			break
		}
	}
	totalChanges := (sizeS - countSame) + (sizeT - countSame)

	if totalChanges > k {
		response = "No"
	} else if totalChanges%2 == k%2 {
		response = "Yes"
	} else if sizeS+sizeT-k < 0 {
		response = "Yes"
	} else {
		response = "No"
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	s := readLine(reader)

	t := readLine(reader)

	kTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := appendAndDelete(s, t, k)

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
