package question_1031_1040

// 1032. 字符流
// https://leetcode-cn.com/problems/stream-of-characters
// Topics: 字典树

type StreamChecker struct {
}

func Constructor(words []string) StreamChecker {
	return StreamChecker{}
}

func (this *StreamChecker) Query(letter byte) bool {
	return false
}

/**
 * Your StreamChecker object will be instantiated and called as such:
 * obj := Constructor(words);
 * param_1 := obj.Query(letter);
 */
