package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var _ = strconv.Itoa // Ignore this comment. You can still use the package "strconv".

	var i uint64 = 4
	var d = 4.0
	var s = "HackerRank "

	scanner := bufio.NewScanner(os.Stdin)
	var data []string
	for scanner.Scan() {
		data = append(data, scanner.Text()) // Println will add back the final '\n'
	}
	a, _ := strconv.ParseInt(data[0], 10, 64)
	b, _ := strconv.ParseFloat(data[1], 64)
	fmt.Printf("%d\n", uint64(a)+i)
	fmt.Printf("%.1f\n", b+d)
	fmt.Printf("%s%s\n", s, data[2])
}
