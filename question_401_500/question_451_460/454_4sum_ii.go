package question_451_460

import (
	"fmt"
)

// 454. 四数相加 II
// https://leetcode-cn.com/problems/4sum-ii
// Topics: 哈希表 二分查找

func fourSumCount(A []int, B []int, C []int, D []int) int {
	var res, flag = 0, make(map[int]int, len(A)*len(B))
	for _, a := range A {
		for _, b := range B {
			flag[a+b]++
		}
	}
	for _, c := range C {
		for _, d := range D {
			res += flag[0-c-d]
		}
	}
	return res
}

func fourSumCount2(A []int, B []int, C []int, D []int) int {
	fmt.Println(len(A))
	var res = 0
	for _, a := range A {
		for _, b := range B {
			for _, c := range C {
				for _, d := range D {
					if a+b+c+d == 0 {
						res++
					}
				}
			}
		}
	}
	return res
}
