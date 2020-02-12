package question_811_820

// 819. 最常见的单词
// https://leetcode-cn.com/problems/most-common-word
// Topics: 字符串

func mostCommonWord(paragraph string, banned []string) string {
	var bannedMap = make(map[string]struct{}, len(banned))
	for _, word := range banned {
		bannedMap[word] = struct{}{}
	}
	paragraph += " "
	start := -1
	var resMap = make(map[string]int)
	for i := 0; i < len(paragraph); i++ {
		if start == -1 && paragraph[i] != ' ' && paragraph[i] != '!' && paragraph[i] != '?' &&
			paragraph[i] != '\'' && paragraph[i] != ',' && paragraph[i] != ';' && paragraph[i] != '.' {
			start = i
		}
		if start != -1 && (paragraph[i] == ' ' || paragraph[i] == '!' || paragraph[i] == '?' ||
			paragraph[i] == '\'' || paragraph[i] == ',' || paragraph[i] == ';' || paragraph[i] == '.') {
			var tmp = make([]byte, i-start)
			for j := start; j < i; j++ {
				if paragraph[j] >= 'A' && paragraph[j] <= 'Z' {
					tmp[j-start] = paragraph[j] - 'A' + 'a'
				} else {
					tmp[j-start] = paragraph[j]
				}
			}
			if _, ok := bannedMap[string(tmp)]; !ok {
				resMap[string(tmp)]++
			}
			start = -1
		}
	}
	var res = ""
	var max = 0
	for key, num := range resMap {
		if num > max {
			max = num
			res = key
		}
	}
	return res
}
