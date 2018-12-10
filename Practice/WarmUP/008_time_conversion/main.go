package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

/*
 * Complete the timeConversion function below.
 */
func timeConversion(s string) string {
	dateId := s[len(s)-2:]
	firstHour, _ := strconv.ParseInt(s[:2], 10, 8)
	if dateId == "AM" {
		if firstHour != 12 {
			return s[:len(s)-2]
		} else {
			return fmt.Sprintf("00%s", s[2:len(s)-2])
		}
	} else if dateId == "PM" {
		if firstHour == 12 {
			return s[:len(s)-2]
		} else {
			return fmt.Sprintf("%d%s", firstHour+12, s[2:len(s)-2])
		}
	} else {
		return ""
	}
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	outputFile, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer outputFile.Close()

	writer := bufio.NewWriterSize(outputFile, 1024*1024)

	s := readLine(reader)

	result := timeConversion(s)

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
