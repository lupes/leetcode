package question_1471_1480

import (
	"sort"
)

// 1451. 重新排列句子中的单词
// https://leetcode-cn.com/problems/rearrange-words-in-a-sentence/
// Topics: 排序 字符串

func getStrongest(arr []int, k int) []int {
	if k == len(arr) {
		return arr
	}

	sort.Ints(arr)

	var res = make([]int, 0, k)
	left, right, middle := 0, len(arr)-1, arr[(len(arr)-1)/2]
	for i := 0; i < k; i++ {
		a := arr[left] - middle
		if a < 0 {
			a *= -1
		}
		b := arr[right] - middle
		if b < 0 {
			b *= -1
		}

		if a > b || (a == b && arr[left] > arr[right]) {
			res = append(res, arr[left])
			left++
		} else {
			res = append(res, arr[right])
			right--
		}
	}

	return res
}
