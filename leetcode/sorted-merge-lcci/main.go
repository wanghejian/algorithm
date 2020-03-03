package main

import (
	"fmt"
)

func merge(A []int, m int, B []int, n int) {
	c := 0
	mm := m
	for i, v := range B {
		for {
			if c == m {
				for j := c; j < mm+n; j++ {
					A[j] = B[i]
					i++
				}
				return
			}

			if A[c] > v {
				for j := m; j > c; j-- {
					A[j] = A[j-1]
				}
				A[c] = v
				c++
				m++
				break
			}
			c++
		}
	}
}

func main() {
	var A [6]int = [6]int{1, 2, 3, 0, 0, 0} // 1 2 2 3 0 0
	var B [3]int = [3]int{4, 5, 6}          // 5 6
	merge(A[:], 3, B[:], 3)
	fmt.Println(A)

	var C [2]int = [2]int{2, 0}
	var F [1]int = [1]int{1}
	merge(C[:], 1, F[:], 1)
	fmt.Println(C)

	var D [1]int = [1]int{10}
	var E [1]int = [1]int{0}
	merge(E[:], 0, D[:], 1)
	fmt.Println(E)
}
