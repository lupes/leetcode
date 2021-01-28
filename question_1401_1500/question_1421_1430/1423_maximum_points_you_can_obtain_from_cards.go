package question_1421_1430

// 1423. 可获得的最大点数
// https://leetcode-cn.com/problems/maximum-points-you-can-obtain-from-cards/
// Topics: 动态规划 滑动窗口

func maxScore(cardPoints []int, k int) int {
	l, sum, min, tmp := len(cardPoints)-k, 0, 10000000000, 0
	for i, n := range cardPoints {
		sum, tmp = sum+n, tmp+n
		if i >= l {
			tmp -= cardPoints[i-l]
		}

		if i >= l-1 && tmp < min {
			min = tmp
		}
	}

	return sum - min
}
