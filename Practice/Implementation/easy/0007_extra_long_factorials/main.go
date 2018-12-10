package main

import (
	"bufio"
	"fmt"
	"io"
	"math/big"
	"os"
	"strconv"
	"strings"
)

// Complete the extraLongFactorials function below.
func extraLongFactorials(n *big.Int) (result *big.Int) {
	result = new(big.Int)
	if n.Cmp(&big.Int{}) == 0 || n.Cmp(&big.Int{}) == -1 {
		result.SetInt64(1)
	} else {
		result.Set(n)
		result.Mul(result, extraLongFactorials(n.Sub(n, big.NewInt(1))))
	}
	return
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 1024*1024)

	nTemp, err := strconv.ParseInt(readLine(reader), 10, 64)
	checkError(err)
	n := big.NewInt(nTemp)

	fact := extraLongFactorials(n)
	fmt.Printf("%d", fact)
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
