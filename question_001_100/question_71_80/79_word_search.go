package question_71_80

// 79. 单词搜索
// https://leetcode-cn.com/problems/word-search
// Topics: 数组 回溯算法

func exist(board [][]byte, word string) bool {
	for r, row := range board {
		for c, b := range row {
			if b == word[0] {
				if existDfs(board, r, c, len(board), len(row), 1, word) {
					return true
				}
			}
		}
	}
	return false
}

func existDfs(board [][]byte, r, c, lr, lc, i int, word string) (e bool) {
	if i == len(word) {
		return true
	}
	var t byte = 0
	t, board[r][c] = board[r][c], t
	defer func() {
		t, board[r][c] = board[r][c], t
	}()
	if (r+1 < lr) && (board[r+1][c] == word[i]) {
		e = e || existDfs(board, r+1, c, lr, lc, i+1, word)
	}
	if !e && (r-1 >= 0) && (board[r-1][c] == word[i]) {
		e = e || existDfs(board, r-1, c, lr, lc, i+1, word)
	}
	if !e && (c+1 < lc) && (board[r][c+1] == word[i]) {
		e = e || existDfs(board, r, c+1, lr, lc, i+1, word)
	}
	if !e && (c-1 >= 0) && (board[r][c-1] == word[i]) {
		e = e || existDfs(board, r, c-1, lr, lc, i+1, word)
	}
	return e
}
