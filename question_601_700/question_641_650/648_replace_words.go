package question_641_650

// 648. 单词替换
// https://leetcode-cn.com/problems/replace-words
// Topics: 字典树 哈希表

func replaceWords(dict []string, sentence string) string {
	var flag = make(map[string]struct{})
	for _, c := range dict {
		flag[c] = struct{}{}
	}
	var res, start = "", 0
	for i := 0; i <= len(sentence); i++ {
		if i == len(sentence) || sentence[i] == ' ' {
			res += sentence[start:i] + " "
			start = i + 1
		} else {
			if _, ok := flag[sentence[start:i]]; ok {
				res += sentence[start:i] + " "
				for i = i + 1; i < len(sentence) && sentence[i] != ' '; i++ {
				}
				start = i + 1
			}
		}
	}
	return res[:len(res)-1]
}
