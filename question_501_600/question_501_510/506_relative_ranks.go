package question_501_510

import (
	"sort"
	"strconv"
)

// 506. 相对名次
// https://leetcode-cn.com/problems/relative-ranks/
// Topics:

func findRelativeRanks(nums []int) []string {
	var tmp = make([]int, len(nums))
	copy(tmp, nums)
	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i] > tmp[j]
	})
	var res []string
	for _, n := range nums {
		for i, t := range tmp {
			if t == n {
				switch i {
				case 0:
					res = append(res, "Gold Medal")
				case 1:
					res = append(res, "Silver Medal")
				case 2:
					res = append(res, "Bronze Medal")
				default:
					res = append(res, strconv.Itoa(i+1))
				}
				break
			}
		}
	}
	return res
}
