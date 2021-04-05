package question_1591_1600

import (
	"strings"
)

// 1592. 重新排列单词间的空格
// https://leetcode-cn.com/problems/rearrange-spaces-between-words/
// Topics: 字符串

func reorderSpaces(text string) string {
	tmp := strings.Split(text, " ")
	spaces := len(tmp) - 1
	var arr []string
	for _, s := range tmp {
		if s != "" {
			arr = append(arr, s)
		}
	}
	if spaces == 0 {
		return text
	}
	if len(arr) == 1 {
		return arr[0] + strings.Repeat(" ", spaces)
	}
	return strings.Join(arr, strings.Repeat(" ", spaces/(len(arr)-1))) + strings.Repeat(" ", spaces%(len(arr)-1))
}
