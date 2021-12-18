package question_411_420

// 419. 甲板上的战舰
// https://leetcode-cn.com/problems/battleships-in-a-board
// Topics:

func countBattleships(board [][]byte) int {
	var res = 0
	for i, row := range board {
		for j, c := range row {
			if c == 'X' && (i == 0 || (i > 0 && board[i-1][j] == '.')) && (j == 0 || (j > 0 && board[i][j-1] == '.')) {
				res++
			}
		}
	}
	return res
}
