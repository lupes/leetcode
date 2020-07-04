package question_411_420

import (
	"fmt"
)

// 414. 第三大的数
// https://leetcode-cn.com/problems/third-maximum-number
// Topics: 数组

func thirdMax(nums []int) int {
	var first, second, third, min = nums[0], 0, 0, nums[0]
	for _, n := range nums {
		if n > first {
			first = n
		}
		if n < min {
			min = n
		}
	}
	fmt.Println(first, min)
	second, third = min, min
	for _, n := range nums {
		if n != first && n > second {
			second = n
		}
	}
	fmt.Println(second, min)
	for _, n := range nums {
		if n != first && n != second && n > third {
			third = n
		}
	}
	if first == second || second == third {
		return first
	} else {
		return third
	}
}
