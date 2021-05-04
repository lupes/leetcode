package question_1661_1670

// 1668. 最大重复子字符串
// https://leetcode-cn.com/problems/maximum-repeating-substring/
// Topics: 字符串

func maxRepeating(sequence string, word string) int {
	l, max := len(word), 0
	for i := range sequence {
		j := 0
		for ; j*l+l+i <= len(sequence) && sequence[i+j*l:i+j*l+l] == word; j++ {
		}
		if j > max {
			max = j
		}
	}
	return max
}
