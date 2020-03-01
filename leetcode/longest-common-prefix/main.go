package main

import (
	"fmt"
	"math"
)

func longestCommonPrefix(strs []string) string {
	min := math.MaxInt64
	minStr := ""
	for _, s := range strs {
		if min > len(s) {
			min = len(s)
			minStr = s
		}
	}

	if min == 0 {
		return ""
	}

	index := -1
	for i, s := range minStr {
		for _, ss := range strs {
			if byte(s) != ss[i] {
				goto OUTER
			}
		}
		index = i
	}

OUTER:
	return minStr[:index+1]
}

func main() {
	fmt.Println(longestCommonPrefix([]string{"flower", "flow", "flight"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "racecar", "car"}))
	fmt.Println(longestCommonPrefix([]string{"dog", "", ""}))
	fmt.Println(longestCommonPrefix([]string{"dog"}))
}
