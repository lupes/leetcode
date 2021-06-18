package question_381_390

import (
	"strings"
)

// 388. 文件的最长绝对路径
// https://leetcode-cn.com/problems/longest-absolute-file-path
// Topics: 栈

func lengthLongestPath(input string) int {
	var res = 0
	arr := strings.Split(input, "\n")
	var stack []string
	for _, file := range arr {
		var i = 0
		for ; file[:1] == "\t"; i++ {
			file = file[1:]
		}
		if i < len(stack) {
			stack = stack[:i]
		}
		stack = append(stack, file)
		if strings.Contains(file, ".") {
			tmp := 0
			for _, d := range stack {
				tmp += len(d) + 1
			}
			if tmp-1 > res {
				res = tmp - 1
			}
		}
	}
	return res
}
