/*
URL			: https://www.hackerrank.com/challenges/separate-the-numbers/problem
AUTHOR		: DmitriyH
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

func numberDigitOf(number int64) int {
	if number/10 == 0 {
		return 1
	} else {
		return 1 + numberDigitOf(number/10)
	}
}

// Complete the separateNumbers function below.
func separateNumbers(s string) string {
	size := len(s)
	if string(s[0]) == "0" {
		return "NO"
	}
	var i int
	for i = 1; i <= size; i++ {
		var (
			firstNumber, _      = strconv.ParseInt(string(s[0:i]), 10, 64)
			numberDigitsCurrent = numberDigitOf(firstNumber)
			numberDigitsNext    = numberDigitOf(firstNumber + 1)
			flag                = true
		)
		if numberDigitsCurrent < size/2+1 {
			if string(s[numberDigitsCurrent]) != "0" {
				w := numberDigitsCurrent
				for w < size {
					firstNumber++
					numberDigitsNext = numberDigitOf(firstNumber)
					limInf := numberDigitsNext+w
					if limInf <= size {
						seqNumber, _ := strconv.ParseInt(string(s[w:limInf]), 10, 64)
						if firstNumber != seqNumber {
							flag = false
							break
						}
					} else {
						return "NO"
					}
					w += numberDigitsNext
				}
				if flag {
					break
				}
			}
		} else {
			return "NO"
		}
	}
	return fmt.Sprintf("%s %s", "YES", string(s[0:i]))
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	qTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	q := int32(qTemp)

	for qItr := 0; qItr < int(q); qItr++ {
		s := readLine(reader)

		fmt.Println(separateNumbers(s))
	}
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
