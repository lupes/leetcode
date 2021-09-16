package question_211_220

// 212. 单词搜索 II
// https://leetcode-cn.com/problems/word-search-ii
// Topics: 字典树 回溯算法

func findWords(board [][]byte, words []string) []string {
	var flag = make(map[string]bool)
	for _, word := range words {
		flag[word] = true
		for i := 1; i < len(word); i++ {
			if _, ok := flag[word[:i]]; !ok {
				flag[word[:i]] = false
			}
		}
	}
	var res []string
	for i, row := range board {
		for j, col := range row {
			prefix := string([]byte{col})
			if e, ok := flag[prefix]; ok {
				if e {
					res = append(res, prefix)
					flag[prefix] = false
				}
				tmp := board[i][j]
				board[i][j] = 0
				res = append(res, findWordsHelper(board, flag, prefix, i, j)...)
				board[i][j] = tmp
			}
		}
	}
	return res
}

func findWordsHelper(board [][]byte, words map[string]bool, prefix string, i, j int) []string {
	var res []string
	for _, t := range [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
		a, b := i+t[0], j+t[1]
		if a >= 0 && a < len(board) && b >= 0 && b < len(board[0]) && board[a][b] != 0 {
			pre := prefix + string([]byte{board[a][b]})
			if e, ok := words[pre]; ok {
				if e {
					res = append(res, pre)
					words[pre] = false
				}
				tmp := board[i][j]
				board[i][j] = 0
				res = append(res, findWordsHelper(board, words, pre, a, b)...)
				board[i][j] = tmp
			}
		}
	}
	return res
}
