package main

import (
	"fmt"
	"math"
)

func reverse(x int) int {
	if x < 10 && x > -10 {
		return x
	}

	y := 0
	for {
		t := x % 10
		x = x / 10
		y = y*10 + t
		if x == 0 {
			break
		}
	}

	if y > math.MaxInt32 || y < math.MinInt32 {
		return 0
	}
	return y
}

func main() {
	fmt.Println(reverse(123))
	fmt.Println(reverse(-1))
	fmt.Println(reverse(-114211))
	fmt.Println(math.MaxInt32)
	fmt.Println(reverse(math.MaxInt32))
	fmt.Println(math.MinInt32)
	fmt.Println(reverse(math.MinInt32))

}
