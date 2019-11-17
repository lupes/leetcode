package question_601_610

import (
	"strings"
)

// 609. 在系统中查找重复文件
// https://leetcode-cn.com/problems/find-duplicate-file-in-system/

func findDuplicate(paths []string) [][]string {
	var files = make(map[string][]string)
	var replacer = strings.NewReplacer("(", " ", ")", "")
	for _, path := range paths {
		arr := strings.Split(replacer.Replace(path), " ")
		for i := 1; i < len(arr); i += 2 {
			files[arr[i+1]] = append(files[arr[i+1]], arr[0]+"/"+arr[i])
		}
	}
	var res [][]string
	for _, v := range files {
		if len(v) > 1 {
			res = append(res, v)
		}
	}
	return res
}
