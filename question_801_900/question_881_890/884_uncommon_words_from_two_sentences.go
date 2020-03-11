package question_881_890

// 884. 两句话中的不常见单词
// https://leetcode-cn.com/problems/uncommon-words-from-two-sentences
// Topics: 哈希表

func uncommonFromSentences(A string, B string) []string {
	var AFlag = uncommonFromSentencesHelper(A)
	var BFlag = uncommonFromSentencesHelper(B)
	var res []string
	for k, v := range AFlag {
		_, ok := BFlag[k]
		if v == 1 && !ok {
			res = append(res, k)
		}
	}
	for k, v := range BFlag {
		_, ok := AFlag[k]
		if v == 1 && !ok {
			res = append(res, k)
		}
	}
	return res
}

func uncommonFromSentencesHelper(A string) map[string]int {
	A += " "
	var start = -1
	var flag = make(map[string]int)
	for i, c := range A {
		if c == ' ' && start != -1 {
			flag[A[start:i]]++
			start = -1
		} else if start == -1 && c != ' ' {
			start = i
		} else if c == ' ' {
			start = -1
		}
	}
	return flag
}
