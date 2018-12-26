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

const (
	negAns = "NO"
	affAns = "YES"
)

func numberDigitOf(number int64) int {
	if number/10 == 0 {
		return 1
	} else {
		return 1 + numberDigitOf(number/10)
	}
}

// Complete the separateNumbers function below.
func separateNumbers(s string) (answer string) {
	answer = negAns
	size := len(s)
	if string(s[0]) == "0" {
		return
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
					lowLim := numberDigitsNext + w

					// Verify if the lower limit if more than size array
					if lowLim <= size {
						seqNumber, _ := strconv.ParseInt(string(s[w:lowLim]), 10, 64)
						if firstNumber != seqNumber {
							flag = false
							break
						}
					} else {
						return
					}
					w += numberDigitsNext
				}
				if flag {
					break
				}
			}
		} else {
			return
		}
	}
	return fmt.Sprintf("%s %s", affAns, string(s[0:i]))
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
