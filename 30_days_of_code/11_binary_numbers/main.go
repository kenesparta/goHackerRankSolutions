package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strconv"
	"strings"
)

func toBin(n int32) (numOnes int32) {
	//var maxOnes int32
	var countOnes int32
	var piv = n
	for piv > 0 {
		// Counting the ones
		if piv%2 == 1 {
			countOnes++
		} else {
			if numOnes <= countOnes {
				numOnes = countOnes
			}
			countOnes = 0
		}

		if piv == 1 && numOnes <= countOnes {
			numOnes = countOnes
		}

		piv = int32(piv / 2)
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := int32(nTemp)
	fmt.Println(toBin(n))
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
