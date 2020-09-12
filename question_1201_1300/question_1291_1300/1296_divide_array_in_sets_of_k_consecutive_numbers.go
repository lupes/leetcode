package question_1291_1300

import (
	"sort"
)

// 1296. 划分数组为连续数字的集合
// https://leetcode-cn.com/problems/divide-array-in-sets-of-k-consecutive-numbers/
// Topics: 贪心算法 数组

func isPossibleDivide(nums []int, k int) bool {
	if len(nums)%k != 0 {
		return false
	}
	var flagMap = make(map[int]int)
	var flagArr = make([]int, 0)
	for _, n := range nums {
		if _, ok := flagMap[n]; !ok {
			flagArr = append(flagArr, n)
		}
		flagMap[n]++
	}
	sort.Ints(flagArr)
	for len(flagArr) > 0 {
		start := flagArr[0]
		if k > len(flagArr) {
			return false
		}
		for i := 0; i < k; i++ {
			if flagArr[i] == start+i {
				flagMap[flagArr[i]]--
				if flagMap[flagArr[i]] == 0 {
					flagArr[i] = 0
				}
			} else {
				return false
			}
		}
		for len(flagArr) > 0 && flagArr[0] == 0 {
			flagArr = flagArr[1:]
		}
	}
	return true
}
