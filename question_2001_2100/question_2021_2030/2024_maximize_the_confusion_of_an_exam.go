package question_2021_2030

// 2024. 考试的最大困扰度
// https://leetcode-cn.com/problems/maximize-the-confusion-of-an-exam/
// Topic: 字符串 二分查找 前缀和 滑动窗口

func maxConsecutiveAnswers(answerKey string, k int) int {
	max1 := maxConsecutiveAnswersHelper(answerKey, k, 'T')
	max2 := maxConsecutiveAnswersHelper(answerKey, k, 'F')
	if max1 > max2 {
		return max1
	}
	return max2
}

func maxConsecutiveAnswersHelper(answerKey string, k int, b int32) int {
	var tmp = make([]int, len(answerKey)+1)
	for i, c := range answerKey {
		if c == b {
			tmp[i+1] = tmp[i] + 1
		} else {
			tmp[i+1] = tmp[i]
		}
	}
	var max = 0
	for l, r := 0, 0; r < len(answerKey)+1; {
		if tmp[r]-tmp[l] <= k {
			if r-l > max {
				max = r - l
			}
			r++
		} else {
			l++
		}
	}
	return max
}
