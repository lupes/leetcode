package question_501_510

import (
	"sort"
	"strconv"
)

// 506. 相对名次
// https://leetcode-cn.com/problems/relative-ranks/
// Topics: 数组 排序 堆

func findRelativeRanks(nums []int) []string {
	var flag = make([][2]int, len(nums))
	for i, n := range nums {
		flag[i] = [2]int{i, n}
	}

	sort.Slice(flag, func(i, j int) bool {
		return flag[i][1] > flag[j][1]
	})

	var res = make([]string, len(nums))
	for i, n := range flag {
		switch i {
		case 0:
			res[n[0]] = "Gold Medal"
		case 1:
			res[n[0]] = "Silver Medal"
		case 2:
			res[n[0]] = "Bronze Medal"
		default:
			res[n[0]] = strconv.Itoa(i + 1)
		}
	}
	return res
}
