package question_1811_1820

// 1816. 截断句子
// https://leetcode-cn.com/problems/truncate-sentence/
// 数组 字符串

func truncateSentence(s string, k int) string {
	for i, c := range s {
		if c == ' ' {
			k--
		}
		if k == 0 {
			return s[:i]
		}
	}
	return s
}
