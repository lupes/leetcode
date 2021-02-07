package question_1451_1460

// 1451. 重新排列句子中的单词
// https://leetcode-cn.com/problems/rearrange-words-in-a-sentence/
// Topics: 排序 字符串

import (
	"sort"
	"strings"
)

func arrangeWords(text string) string {
	arr := strings.Split(text, " ")

	a0 := []byte(arr[0])
	a0[0] += 'a' - 'A'
	arr[0] = string(a0)

	sort.SliceStable(arr, func(i, j int) bool {
		return len(arr[i]) < len(arr[j])
	})

	a0 = []byte(arr[0])
	a0[0] -= 'a' - 'A'
	arr[0] = string(a0)

	return strings.Join(arr, " ")
}
