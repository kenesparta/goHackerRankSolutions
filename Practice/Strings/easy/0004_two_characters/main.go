/*
URL			: https://www.hackerrank.com/challenges/two-characters/problem
AUTHOR		: nabila_ahmed
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

// Verify if the letter is not in the array
func notIn(v string, arr []string) bool {
	for _, nr := range arr {
		if v == nr {
			return false
		}
	}
	return true
}

// Adjacent repeated characters true if exist, false if not.
func adjRepChar(s string) bool {
	for i := 0; i < len(s)-1; i++ {
		if s[i] == s[i+1] {
			return true
		}
	}
	return false
}

type TwoString struct {
	PermLetters [][]string
	RepLetter   []string
	MainString  string
}

// Init
func (ts *TwoString) Initialize(s string) {
	ts.MainString = s
	ts.RepLetter = ts.searchRepLetter()
	ts.notRepeatPerm()
}

// Not repeat permutation
func (ts *TwoString) notRepeatPerm() {
	for i, nr := range ts.RepLetter {
		for _, v := range ts.RepLetter[i+1:] {
			var resultString []string
			resultString = append(resultString, nr)
			resultString = append(resultString, v)
			ts.PermLetters = append(ts.PermLetters, resultString)
		}
	}
	return
}

// Flter the array of strings by the conditions and return the max lengt of these
func (ts *TwoString) filterByConditions() int32 {
	var (
		posString = ts.PossibleStrings()
		maxLengt  int32
	)

	for _, v := range posString {
		if !adjRepChar(v) {
			size := int32(len(v))
			if maxLengt < size {
				maxLengt = size
			}
		}
	}
	return maxLengt
}

// Search all not repeated letters on the string
func (ts *TwoString) searchRepLetter() []string {
	firstLetter := string(ts.MainString[0])
	notRepeat := []string{firstLetter}
	for _, v := range ts.MainString {
		v_str := string(v)
		if notIn(v_str, notRepeat) {
			notRepeat = append(notRepeat, v_str)
		}
	}
	return notRepeat
}

// Possible strings to be accepted
func (ts *TwoString) PossibleStrings() (posString []string) {
	for _, nr := range ts.PermLetters {
		resultString := ""
		for _, v := range ts.MainString {
			if string(v) == nr[0] || string(v) == nr[1] {
				resultString += string(v)
			}
		}
		posString = append(posString, resultString)
	}
	return
}

// Complete the alternate function below.
func alternate(s string, l int) int32 {
	twoS := &TwoString{}
	twoS.Initialize(s)
	return twoS.filterByConditions()
}

func main() {
	reader := bufio.NewReaderSize(os.Stdin, 16*1024*1024)

	stdout, err := os.Create(os.Getenv("OUTPUT_PATH"))
	checkError(err)

	defer stdout.Close()

	writer := bufio.NewWriterSize(stdout, 16*1024*1024)

	lTemp, err := strconv.ParseInt(strings.TrimSpace(readLine(reader)), 10, 64)
	checkError(err)
	l := int(lTemp)

	s := readLine(reader)

	result := alternate(s, l)

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
