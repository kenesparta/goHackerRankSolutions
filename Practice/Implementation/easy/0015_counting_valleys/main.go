package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

// Complete the countingValleys function below.
func countingValleys(n int32, s string) (numberValley int32) {
	var hikeSum int32
	var hikeStatus string
	// 0: Sea Level ; M: Mountain; V: Valley
	hikeStatusAnt := "0"
	for _, v := range s {
		if string(v) == "U" {
			hikeSum++
		} else if string(v) == "D" {
			hikeSum--
		}

		if hikeSum > 0 {
			hikeStatus = "M"
		} else if hikeSum < 0 {
			hikeStatus = "V"
		} else {
			hikeStatus = "0"
		}

		if hikeStatusAnt == "0" && hikeStatus == "V" {
			numberValley++
		}

		hikeStatusAnt = hikeStatus
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

	result := countingValleys(n, s)

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
