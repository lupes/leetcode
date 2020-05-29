package question_1131_1140

import (
	"strings"
)

// 1138. 字母板上的路径
// https://leetcode-cn.com/problems/alphabet-board-path
// Topics: 哈希表 字符串

var f = [26][2]int{
	{0, 0}, {0, 1}, {0, 2}, {0, 3}, {0, 4},
	{1, 0}, {1, 1}, {1, 2}, {1, 3}, {1, 4},
	{2, 0}, {2, 1}, {2, 2}, {2, 3}, {2, 4},
	{3, 0}, {3, 1}, {3, 2}, {3, 3}, {3, 4},
	{4, 0}, {4, 1}, {4, 2}, {4, 3}, {4, 4},
	{5, 0},
}

func alphabetBoardPath(target string) string {
	var now, res = [2]int{0, 0}, ""
	for _, c := range target {
		if now[0] == 5 && c != 'z' {
			now = [2]int{4, 0}
			res += "U"
		}
		t := f[c-'a']
		x, a := "D", t[0]-now[0]
		if a < 0 {
			x, a = "U", -a
		}

		y, b := "R", t[1]-now[1]
		if b < 0 {
			y, b = "L", -b
		}
		res += strings.Repeat(y, b) + strings.Repeat(x, a) + "!"
		now = t
	}
	return res
}
