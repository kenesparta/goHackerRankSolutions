package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
)

func expectedSignal(n int32) (signal string) {
	for i := int32(0); i < n; i++ {
		signal += "SOS"
	}
	return
}

// Complete the marsExploration function below.
func marsExploration(s string) (errors int32) {
	size := int32(len(s))
	expSignal := expectedSignal(size / 3)
	for i := int32(0); i < size; i++ {
		if expSignal[i] != s[i] {
			errors++
		}
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

	result := marsExploration(s)

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
