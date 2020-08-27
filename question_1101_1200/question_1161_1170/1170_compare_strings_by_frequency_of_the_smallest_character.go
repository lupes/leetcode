package question_1161_1170

// 1170. 比较字符串最小字母出现频次
// https://leetcode-cn.com/problems/compare-strings-by-frequency-of-the-smallest-character
// Topics: 数组 字符串

func numSmallerByFrequency(queries []string, words []string) []int {
	flag := make([]int, 12)
	var res = make([]int, len(queries))
	for _, word := range words {
		flag[numSmallerByFrequencyHelper(word)]++
	}
	for i := 10; i >= 0; i-- {
		flag[i] += flag[i+1]
	}
	for i, query := range queries {
		n := numSmallerByFrequencyHelper(query)
		res[i] = flag[n+1]
	}
	return res
}

func numSmallerByFrequencyHelper(word string) int {
	var f = make([]int, 26)
	for _, c := range word {
		f[c-'a']++
	}
	for _, n := range f {
		if n > 0 {
			return n
		}
	}
	return 0
}
