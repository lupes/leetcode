package question_71_80

import "fmt"

// 72. 编辑距离
// https://leetcode-cn.com/problems/edit-distance/
// Topics: 字符串 动态规划

func Print(word1, word2 string, flag [][]int) {
	word1, word2 = " -"+word1, "-"+word2
	for _, c := range word1 {
		fmt.Printf(" %c", c)
	}
	fmt.Println()
	for i, v := range flag {
		fmt.Printf("%c %+v\n", word2[i], v)
	}
	fmt.Println()
}

func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	flag := make([][]int, l2+1)
	for i := range flag {
		flag[i] = make([]int, l1+1)
		flag[i][0] = i
	}
	for i := range flag[0] {
		flag[0][i] = i
	}
	for i, w2 := range word2 {
		for j, w1 := range word1 {
			if w2 == w1 {
				flag[i+1][j+1] = flag[i][j]
			} else {
				t := flag[i][j]
				if flag[i+1][j] < t {
					t = flag[i+1][j]
				}
				if flag[i][j+1] < t {
					t = flag[i][j+1]
				}
				flag[i+1][j+1] = t + 1
			}
		}
	}
	Print(word1, word2, flag)
	return flag[l2][l1]
}
