package question_1471_1480

// 1476. 子矩形查询
// https://leetcode-cn.com/problems/subrectangle-queries/
// Topics: 设计

type SubrectangleQueries struct {
	rectangle [][]int
}

func Constructor6(rectangle [][]int) SubrectangleQueries {
	return SubrectangleQueries{rectangle: rectangle}
}

func (this *SubrectangleQueries) UpdateSubrectangle(row1 int, col1 int, row2 int, col2 int, newValue int) {
	for i := row1; i <= row2; i++ {
		for j := col1; j <= col2; j++ {
			this.rectangle[i][j] = newValue
		}
	}
}

func (this *SubrectangleQueries) GetValue(row int, col int) int {
	return this.rectangle[row][col]
}
