package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func abs(n int32) (valAbs int32) {
	if n < 0 {
		valAbs = -n
	} else {
		valAbs = n
	}
	return
}

func reverse(n int32) (rev int32) {
	for n > 0 {
		remainder := n % 10
		rev *= 10
		rev += remainder
		n /= 10
	}
	return
}

// Complete the beautifulDays function below.
func beautifulDays(i int32, j int32, k int32) (bDaysCount int32) {
	for idx := i; idx <= j; idx++ {
		if abs(idx-reverse(idx))%k == 0 {
			bDaysCount++
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

	ijk := strings.Split(readLine(reader), " ")

	iTemp, err := strconv.ParseInt(ijk[0], 10, 64)
	checkError(err)
	i := int32(iTemp)

	jTemp, err := strconv.ParseInt(ijk[1], 10, 64)
	checkError(err)
	j := int32(jTemp)

	kTemp, err := strconv.ParseInt(ijk[2], 10, 64)
	checkError(err)
	k := int32(kTemp)

	result := beautifulDays(i, j, k)

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
