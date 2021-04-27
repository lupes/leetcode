package question_1651_1660

// 1657. 确定两个字符串是否接近
// https://leetcode-cn.com/problems/determine-if-two-strings-are-close/
// Topics: 贪心算法

import (
	"sort"
)

func closeStrings(word1 string, word2 string) bool {
	if len(word1) != len(word2) {
		return false
	}
	var flag1, flag2 = make([]int, 26), make([]int, 26)
	for i := range word1 {
		flag1[word1[i]-'a']++
		flag2[word2[i]-'a']++
	}

	for i := range flag1 {
		if !((flag1[i] > 0 && flag2[i] > 0) || flag1[i] == 0 && flag2[i] == 0) {
			return false
		}
	}

	sort.Ints(flag1)
	sort.Ints(flag2)

	for i := range flag1 {
		if flag1[i] != flag2[i] {
			return false
		}
	}

	return true
}
