package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func sumPermutations(S []int32) (sumPerm []int32) {
	var mapMultiplers []int32
	i := 0
	sizeS := len(S)
	for i < sizeS {
		j := i + 1
		for j < sizeS {
			sumPerm = append(sumPerm, S[i]+S[j])
			j++
		}
		i++
	}
	return
}

// Complete the nonDivisibleSubset function below.
func nonDivisibleSubset(k int32, S []int32) int32 {
	fmt.Println(sumPermutations(S))
	return 0
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	nk := strings.Split(readLine(reader), " ")

	nTemp, err := strconv.ParseInt(nk[0], 10, 64)
	checkError(err)
	n := int32(nTemp)

	kTemp, err := strconv.ParseInt(nk[1], 10, 64)
	checkError(err)
	k := int32(kTemp)

	STemp := strings.Split(readLine(reader), " ")

	var S []int32

	for i := 0; i < int(n); i++ {
		SItemTemp, err := strconv.ParseInt(STemp[i], 10, 64)
		checkError(err)
		SItem := int32(SItemTemp)
		S = append(S, SItem)
	}

	result := nonDivisibleSubset(k, S)

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
