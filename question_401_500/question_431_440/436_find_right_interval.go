package question_431_440

import (
	"sort"
)

// 436. 寻找右区间
// https://leetcode-cn.com/problems/find-right-interval
// Topics: 二分查找

func findRightInterval(intervals [][]int) []int {
	var flag, keys = make(map[int]int), make([]int, 0)
	for i, interval := range intervals {
		_, ok := flag[interval[0]]
		if !ok {
			flag[interval[0]] = i
			keys = append(keys, interval[0])
		}
	}
	sort.Ints(keys)
	var res = make([]int, len(intervals))
	for i, interval := range intervals {
		e, n := findRightIntervalHelper(keys, interval[1])
		if e {
			res[i] = flag[n]
		} else {
			res[i] = -1
		}
	}
	return res
}

func findRightIntervalHelper(keys []int, key int) (bool, int) {
	if keys[len(keys)-1] < key {
		return false, 0
	}
	l, r := 0, len(keys)
	for r > l {
		c := (l + r) / 2
		if keys[c] > key {
			r = c
		} else if keys[c] == key {
			return true, key
		} else {
			l = c + 1
		}
	}
	return true, keys[l]
}
