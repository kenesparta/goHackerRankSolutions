/*
URL			: https://www.hackerrank.com/challenges/halloween-sale/problem
AUTHOR		: kevinsogo
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

// Complete the howManyGames function below.
func howManyGames(p int32, d int32, m int32, s int32) (numberGames int32) {
	// Return the number of games you can buy
	totalSum := p
	for totalSum <= s {
		p -= d
		if p < m {
			p = m
		}
		numberGames++
		totalSum += p
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	pdms := strings.Split(readLine(reader), " ")

	pTemp, err := strconv.ParseInt(pdms[0], 10, 64)
	checkError(err)
	p := int32(pTemp)

	dTemp, err := strconv.ParseInt(pdms[1], 10, 64)
	checkError(err)
	d := int32(dTemp)

	mTemp, err := strconv.ParseInt(pdms[2], 10, 64)
	checkError(err)
	m := int32(mTemp)

	sTemp, err := strconv.ParseInt(pdms[3], 10, 64)
	checkError(err)
	s := int32(sTemp)

	answer := howManyGames(p, d, m, s)

	fmt.Fprintf(writer, "%d\n", answer)

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
