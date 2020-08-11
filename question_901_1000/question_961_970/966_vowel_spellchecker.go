package question_961_970

// 966. 元音拼写检查器
// https://leetcode-cn.com/problems/vowel-spellchecker
// Topics: 哈希表 字符串

func tof2(c byte) byte {
	if c < 'a' {
		return c + 32
	} else {
		return c
	}
}

func tof3(c byte) byte {
	if c == 'a' || c == 'A' || c == 'e' || c == 'E' || c == 'i' || c == 'I' || c == 'o' || c == 'O' || c == 'u' || c == 'U' {
		return 'a'
	} else if c < 'a' {
		return c + 32
	} else {
		return c
	}
}

func spellchecker(wordlist []string, queries []string) []string {
	var flag1 = make(map[string]struct{})
	var flag2 = make(map[string]int)
	var flag3 = make(map[string]int)
	var res = make([]string, len(queries))
	for i, word := range wordlist {
		f2, f3 := make([]byte, len(word)), make([]byte, len(word))
		for i, c := range word {
			f2[i] = tof2(byte(c))
			f3[i] = tof3(byte(c))
		}
		if _, ok := flag1[word]; !ok {
			flag1[word] = struct{}{}
		}
		if _, ok := flag2[string(f2)]; !ok {
			flag2[string(f2)] = i
		}
		if _, ok := flag3[string(f3)]; !ok {
			flag3[string(f3)] = i
		}
	}
	for i, query := range queries {
		if _, ok := flag1[query]; ok {
			res[i] = query
			continue
		}
		f2, f3 := make([]byte, len(query)), make([]byte, len(query))
		for i, c := range query {
			f2[i] = tof2(byte(c))
			f3[i] = tof3(byte(c))
		}
		if j, ok := flag2[string(f2)]; ok {
			res[i] = wordlist[j]
			continue
		}
		if j, ok := flag3[string(f3)]; ok {
			res[i] = wordlist[j]
			continue
		}
	}
	return res
}
