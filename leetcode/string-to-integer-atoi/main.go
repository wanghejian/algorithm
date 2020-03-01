package main

import (
	"fmt"
	"math"
)

func myAtoi(str string) int {
	var (
		start bool
		num   int
		sign  int = 1
	)

	for _, s := range str {

		if s == ' ' {
			if !start {
				continue
			} else {
				break
			}
		} else if s == '+' {
			if !start {
				start = true
				continue
			} else {
				break
			}
		} else if s == '-' {
			if !start {
				start = true
				sign = -1
				continue
			} else {
				break
			}
		} else if s <= '9' && s >= '0' {
			if !start {
				start = true
			}
			num = num*10 + int(s) - int('0')
		} else if start {
			break
		} else {
			break
		}

		if num*sign > math.MaxInt32 {
			return math.MaxInt32
		} else if num*sign < math.MinInt32 {
			return math.MinInt32
		}
	}

	return num * sign
}

func main() {
	fmt.Println(myAtoi("100"))
	fmt.Println(myAtoi("-100"))
	fmt.Println(myAtoi("  1"))
	fmt.Println(myAtoi("  -234"))
	fmt.Println(myAtoi("a100b100"))
	fmt.Println(myAtoi("10000000000000"))
	fmt.Println(myAtoi("-10000000000000"))
	fmt.Println(myAtoi("words and 987"))
	fmt.Println(myAtoi("   +0 123"))
	fmt.Println(myAtoi("2147483648"))
}
