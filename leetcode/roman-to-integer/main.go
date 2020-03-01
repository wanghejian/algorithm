package main

import (
	"fmt"
)

var dict map[string]int

func init() {
	dict = make(map[string]int)
	dict["I"] = 1
	dict["IV"] = 4
	dict["V"] = 5
	dict["IX"] = 9
	dict["X"] = 10
	dict["XL"] = 40
	dict["L"] = 50
	dict["XC"] = 90
	dict["C"] = 100
	dict["CD"] = 400
	dict["D"] = 500
	dict["CM"] = 900
	dict["M"] = 1000
}

func romanToInt(s string) int {
	var sum, i, h, l int
	for {
		if i <= (len(s) - 2) {
			if v, ok := dict[s[i:i+2]]; ok {
				i += 2
				l = v
			} else if v, ok := dict[s[i:i+1]]; ok {
				i += 1
				l = v
			} else {
				return -1
			}
		} else if i <= (len(s) - 1) {
			if v, ok := dict[s[i:i+1]]; ok {
				i += 1
				l = v
			} else {
				return -2
			}
		} else {
			break
		}

		if h != 0 && h < l {
			return -3
		}

		h = l
		sum += l
	}

	return sum
}

func main() {
	fmt.Println(romanToInt("I"))
	fmt.Println(romanToInt("II"))
	fmt.Println(romanToInt("III"))
	fmt.Println(romanToInt("IIII"))
	fmt.Println(romanToInt("IV"))
	fmt.Println(romanToInt("MCMXCIV"))
}
