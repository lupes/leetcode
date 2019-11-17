package question_71_80

import "strings"

// 71. 简化路径
// https://leetcode-cn.com/problems/simplify-path
// Topics: 栈 字符串

func simplifyPath(path string) string {
	var arr []string
	var t string
	for _, s := range path + "/" {
		if s == '/' && t != "" {
			if t == ".." && len(arr) > 0 {
				arr = arr[:len(arr)-1]
			} else if t != ".." && t != "." {
				arr = append(arr, t)
			}
			t = ""
		}
		if s != '/' {
			t += string(s)
		}
	}
	return "/" + strings.Join(arr, "/")
}
