package question_1451_1460

// 1460. 通过翻转子数组使两个数组相等
// https://leetcode-cn.com/problems/make-two-arrays-equal-by-reversing-sub-arrays/
// Topics: 数组

func canBeEqual(target []int, arr []int) bool {
	var flag = make(map[int]int, len(target)/2)
	for _, n := range arr {
		flag[n]++
	}

	for _, n := range target {
		flag[n]--
		if flag[n] < 0 {
			return false
		}
	}

	for _, n := range flag {
		if n != 0 {
			return false
		}
	}
	return true
}
