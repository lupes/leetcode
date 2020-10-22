package question_1021_1030

import (
	"regexp"
	"strings"
)

// 1023. 驼峰式匹配
// https://leetcode-cn.com/problems/camelcase-matching
// Topics: 字典树 字符串

func camelMatch(queries []string, pattern string) []bool {
	pattern = "^" + pattern + "$"
	var arr = make([]string, 0, len(pattern))
	for i := range pattern {
		arr = append(arr, pattern[i:i+1])
	}
	reg, _ := regexp.Compile(strings.Join(arr, "[a-z]*"))

	var res = make([]bool, 0, len(queries))
	for _, query := range queries {
		res = append(res, reg.MatchString(query))
	}

	return res
}
