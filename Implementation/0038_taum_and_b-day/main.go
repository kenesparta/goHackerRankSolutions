/*
URL			: https://www.hackerrank.com/challenges/taum-and-bday/problem
AUTHOR		: amititkgp
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

func min(a, b int64) int64 {
	if a < b {
		return a
	} else {
		return b
	}
}

// Complete the taumBday function below.
func taumBday(b, w, bc, wc, z int64) int64 {
	cfW := wc + z // Conversion factor white
	cfB := bc + z // Conversion factor black

	if cfB < wc || cfW < bc {
		if wc > bc {
			wc = min(cfB, cfW)
		} else {
			bc = min(cfB, cfW)
		}
	}

	return b*bc + w*wc
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 1024*1024)

	tTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	t := int64(tTemp)

	for tItr := 0; tItr < int(t); tItr++ {
		bw := strings.Split(readLine(reader), " ")

		bTemp, err := strconv.ParseInt(bw[0], 10, 64)
		checkError(err)
		b := int64(bTemp)

		wTemp, err := strconv.ParseInt(bw[1], 10, 64)
		checkError(err)
		w := int64(wTemp)

		bcWcz := strings.Split(readLine(reader), " ")

		bcTemp, err := strconv.ParseInt(bcWcz[0], 10, 64)
		checkError(err)
		bc := int64(bcTemp)

		wcTemp, err := strconv.ParseInt(bcWcz[1], 10, 64)
		checkError(err)
		wc := int64(wcTemp)

		zTemp, err := strconv.ParseInt(bcWcz[2], 10, 64)
		checkError(err)
		z := int64(zTemp)

		result := taumBday(b, w, bc, wc, z)

		fmt.Fprintf(writer, "%d\n", result)
	}

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
