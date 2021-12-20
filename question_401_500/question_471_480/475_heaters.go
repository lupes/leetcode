package question_471_480

import (
	"sort"
)

// 475. 供暖器
// https://leetcode-cn.com/problems/heaters
// Topics: 二分查找

func findRadius(houses []int, heaters []int) int {
	sort.Ints(heaters)
	var max, l = 0, len(heaters)
	for _, h := range houses {
		n := sort.SearchInts(heaters, h)
		if n == 0 {
			if heaters[0]-h > max {
				max = heaters[0] - h
			}
		} else if n == l {
			if h-heaters[l-1] > max {
				max = h - heaters[l-1]
			}
		} else if heaters[n] != h {
			tmp := h - heaters[n-1]
			if heaters[n]-h < h-heaters[n-1] {
				tmp = heaters[n] - h
			}
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}
