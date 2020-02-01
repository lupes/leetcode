package question_731_740

// 733. 图像渲染
// https://leetcode-cn.com/problems/flood-fill
// Topics: 深度优先搜索

func floodFill(image [][]int, sr int, sc int, newColor int) [][]int {
	floodFillHelper(image, sr, sc, newColor, image[sr][sc])
	return image
}

func floodFillHelper(image [][]int, sr int, sc int, newColor, oldColor int) {
	if sr >= len(image) || sr < 0 || sc >= len(image[0]) || sc < 0 || image[sr][sc] != oldColor {
		return
	}
	image[sr][sc] = -1
	floodFillHelper(image, sr-1, sc, newColor, oldColor)
	floodFillHelper(image, sr+1, sc, newColor, oldColor)
	floodFillHelper(image, sr, sc-1, newColor, oldColor)
	floodFillHelper(image, sr, sc+1, newColor, oldColor)
	image[sr][sc] = newColor
}
