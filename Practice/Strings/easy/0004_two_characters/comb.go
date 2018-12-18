package main

import (
	"fmt"
)

func combinationsTesting(k int32, rest, arr []string) (perms [][]string) {
	if k > 0 {
		for i, _ := range arr {
			perms = append(perms, combinationsTesting(k-1, append(rest, arr[i]), arr[i+1:])...)
		}
		return perms
	} else {
		return [][]string{rest}
	}
}

func main() {
	fmt.Println(combinationsTesting(2, []string{}, []string{"a", "s", "d", "c", "b", "g", "f", "n", "h"}))
}
