package main

import (
	"fmt"
)

func isPalindrome1(x int) bool {
	if x < 0 {
		return false
	}

	if x < 10 {
		return true
	}

	ch := gen(x)

	sum := 0
	for v := range ch {
		// fmt.Print(v, " ")
		sum = sum*10 + v
	}
	// fmt.Println()
	// fmt.Println(sum)

	if sum != x {
		return false
	}

	return true
}

func gen(x int) <-chan int {
	ch := make(chan int, 24)
	go func() {
		for {
			t := x % 10
			x /= 10
			ch <- t
			if x == 0 {
				close(ch)
				break
			}
		}
	}()

	return ch
}

// O(log 10(n))
func isPalindrome2(x int) bool {
	if x < 0 || (x%10 == 0 && x != 0) {
		return false
	}

	if x < 10 {
		return true
	}

	o := x
	s := 0
	for o > s {
		t := o % 10
		o /= 10

		s = s*10 + t
	}

	return o == s || o == s/10
}

func main() {
	fmt.Println(isPalindrome2(121))
	fmt.Println(isPalindrome2(-121))
	fmt.Println(isPalindrome2(1))
	fmt.Println(isPalindrome2(1221))
	fmt.Println(isPalindrome2(1122))
	fmt.Println(isPalindrome2(100))

}
