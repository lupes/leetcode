package question_71_80

import "fmt"

// 72. 编辑距离
// https://leetcode-cn.com/problems/edit-distance/

func Print(word1, word2 string, args [][]int) {
	word1, word2 = " -"+word1, "-"+word2
	for _, c := range word1 {
		fmt.Printf(" %c", c)
	}
	fmt.Println()
	for i, arg := range args {
		fmt.Printf("%c %+v\n", word2[i], arg)
	}
	fmt.Println()
}

func minDistance(word1 string, word2 string) int {
	l1, l2 := len(word1), len(word2)
	flags := make([][]int, l2+1)
	for i := range flags {
		flags[i] = make([]int, l1+1)
		flags[i][0] = i
	}
	for i := range flags[0] {
		flags[0][i] = i
	}
	for i, w2 := range word2 {
		for j, w1 := range word1 {
			if w2 == w1 {
				flags[i+1][j+1] = flags[i][j]
			} else {
				t := flags[i][j]
				if flags[i+1][j] < t {
					t = flags[i+1][j]
				}
				if flags[i][j+1] < t {
					t = flags[i][j+1]
				}
				flags[i+1][j+1] = t + 1
			}
		}
	}
	return flags[l2][l1]
}
