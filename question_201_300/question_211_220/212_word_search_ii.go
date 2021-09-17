package question_211_220

// 212. 单词搜索 II
// https://leetcode-cn.com/problems/word-search-ii
// Topics: 字典树 回溯算法

func findWords(board [][]byte, words []string) []string {
	var flag = make(map[string]bool, len(words))
	for _, word := range words {
		for i := 1; i < len(word); i++ {
			flag[word[:i]] = false
		}
	}
	for _, word := range words {
		flag[word] = true
	}
	var res []string
	for i, row := range board {
		for j := range row {
			findWordsHelper(board, flag, &res, nil, i, j)
		}
	}
	return res
}

func findWordsHelper(board [][]byte, words map[string]bool, res *[]string, prefix []byte, i, j int) {
	if i < 0 || i >= len(board) || j < 0 || j >= len(board[0]) || board[i][j] == 0 {
		return
	}

	prefix = append(prefix, board[i][j])

	if has, ok := words[string(prefix)]; !ok {
		return
	} else if has {
		*res = append(*res, string(prefix))
		words[string(prefix)] = false
	}

	tmp := board[i][j]
	board[i][j] = 0

	for _, t := range [][]int{{0, -1}, {0, 1}, {-1, 0}, {1, 0}} {
		findWordsHelper(board, words, res, prefix, i+t[0], j+t[1])
	}
	board[i][j] = tmp
}
