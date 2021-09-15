package question_1861_1870

// 1864. 构成交替字符串需要的最小交换次数
// https://leetcode-cn.com/problems/minimum-number-of-swaps-to-make-the-binary-string-alternating/
// Topics: 贪心 字符串

const MaxSwaps = 100000

func minSwaps(s string) int {
	min := MaxSwaps
	if tmp := minSwapsHelper(s, "10"); min > tmp {
		min = tmp
	}
	if min == 0 {
		return 0
	}

	if tmp := minSwapsHelper(s, "01"); min > tmp {
		min = tmp
	}
	if min == MaxSwaps {
		return -1
	}
	return min
}

func minSwapsHelper(s, order string) int {
	var tmp = []int{0, 0}
	for i := range s {
		if order[i%2] != s[i] {
			tmp[s[i]-'0']++
		}
	}
	if tmp[0] == tmp[1] {
		return tmp[0]
	}
	return MaxSwaps
}
