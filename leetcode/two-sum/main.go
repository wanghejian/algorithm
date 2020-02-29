package main

import (
	"fmt"
)

// o(N2)
func twoSum1(nums []int, target int) []int {
	var rst []int
	for i, v := range nums {
		tmp := target - v

		if i == len(nums) {
			return rst
		}

		for j, vv := range nums[i+1:] {

			if tmp == vv {
				rst = append(rst, i, j+i+1)
				break
			}
		}
	}

	return rst
}

// o(2N)
func twoSum2(nums []int, target int) []int {
	var rst []int
	tmp := make(map[int][]int)
	for i, v := range nums {
		vv := tmp[v]
		vv = append(vv, i)
		tmp[v] = vv
	}

	for i, v := range nums {

		k := tmp[v]
		if len(k) <= 1 {
			delete(tmp, v)
		} else {
			tmp[v] = k[1:]
		}

		val := target - v
		if vv, ok := tmp[val]; ok {
			vvv := vv[0]
			rst = append(rst, i, vvv)
			if len(vv) == 1 {
				delete(tmp, val)
			} else {
				tmp[val] = vv[1:]
			}
		}
	}

	return rst
}

// o(N)
func twoSum3(nums []int, target int) []int {
	m := make(map[int]int)
	var rst []int

	for i, v := range nums {
		tmp := target - v
		if vv, ok := m[tmp]; ok {
			rst = append(rst, i, vv)
			delete(m, tmp)
		} else {
			m[v] = i
		}
	}

	return rst
}

func main() {
	rst := twoSum3([]int{2, 3, 6, 7, 11, 15}, 9)
	fmt.Println(rst)
	rst = twoSum3([]int{3, 2, 4}, 6)
	fmt.Println(rst)
	rst = twoSum3([]int{0, 4, 3, 0}, 0)
	fmt.Println(rst)
	rst = twoSum3([]int{-1, -2, -3, -4, -5}, -8)
	fmt.Println(rst)
}
