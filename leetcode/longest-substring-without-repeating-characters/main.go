package main

import (
	"fmt"
)

// o(N)
func lengthOfLongestSubstring(s string) int {
	if len(s) == 0 {
		return 0
	} else if len(s) == 1 {
		return 1
	}

	max := 0
	m := make(map[rune]int)

	index := 0
	for i, v := range s {
		if vv, ok := m[v]; ok && vv >= index {
			index = vv + 1
		}

		m[v] = i

		if max < i-index+1 {
			max = i - index + 1
		}
	}

	return max
}

func main() {
	fmt.Println(lengthOfLongestSubstring("abcabcbb"))
	fmt.Println(lengthOfLongestSubstring("bbbb"))
	fmt.Println(lengthOfLongestSubstring("pwwkew"))
	fmt.Println(lengthOfLongestSubstring("au"))
	fmt.Println(lengthOfLongestSubstring(""))
	fmt.Println(lengthOfLongestSubstring(" "))
	fmt.Println(lengthOfLongestSubstring("abba"))
}
