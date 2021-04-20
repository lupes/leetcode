package question_1631_1640

// 1636. 按照频率将数组升序排序
// https://leetcode-cn.com/problems/sort-array-by-increasing-frequency/
// Topics: 排序 数组

import (
	"sort"
)

func frequencySort(nums []int) []int {
	var flag = make(map[int]int)
	for _, n := range nums {
		flag[n]++
	}
	var arr [][2]int
	for num, times := range flag {
		arr = append(arr, [2]int{num, times})
	}
	sort.Slice(arr, func(i, j int) bool {
		return arr[i][1] < arr[j][1] || (arr[i][1] == arr[j][1] && arr[i][0] > arr[j][0])
	})

	var res []int
	for _, tmp := range arr {
		num, times := tmp[0], tmp[1]
		for i := 0; i < times; i++ {
			res = append(res, num)
		}
	}

	return res
}
