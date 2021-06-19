package question_311_320

// 318. 最大单词长度乘积
// https://leetcode-cn.com/problems/maximum-product-of-word-lengths
// Topics: 位运算

func maxProduct(words []string) int {
	var res, flag = 0, make([]int, len(words))
	for i, word := range words {
		for _, c := range word {
			flag[i] = flag[i] | 1<<(c-'a')
		}
		for j := 0; j < i; j++ {
			if flag[i]&flag[j] == 0 && len(words[i])*len(words[j]) > res {
				res = len(words[i]) * len(words[j])
			}
		}
	}
	return res
}
