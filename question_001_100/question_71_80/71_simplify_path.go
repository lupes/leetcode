package question_71_80

import "strings"

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
