package question_1561_1570

// 1567. 乘积为正数的最长子数组长度
// https://leetcode-cn.com/problems/maximum-length-of-subarray-with-positive-product/
// Topics: 贪心算法

import (
	"fmt"
)

func getMaxLen(nums []int) int {
	var left, right, max, sum, start, end = 0, 0, 0, 0, 0, 0
	for i, n := range nums {
		if n < 0 {
			sum++
			if sum == 1 {
				left = i
				if left-start > max {
					max = left - start
				}
			} else {
				right = i
			}
		}

		if n == 0 || i == len(nums)-1 {

			end = i

			if n == 0 {
				end = i - 1
			}
			if sum%2 == 1 {
				if left-start > end-right {
					end = right - 1
				} else {
					start = left + 1
				}
			}
			if end-start+1 > max {
				max = end - start + 1
			}
			fmt.Printf("start:%d end:%d sum:%d i:%d\n", start, end, sum, i)
			start, sum = i+1, 0
		}
	}
	return max
}
