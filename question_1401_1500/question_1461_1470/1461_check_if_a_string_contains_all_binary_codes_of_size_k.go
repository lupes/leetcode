package question_1461_1470

// 1451. 重新排列句子中的单词
// https://leetcode-cn.com/problems/rearrange-words-in-a-sentence/
// Topics: 排序 字符串

func hasAllCodes(s string, k int) bool {
	var flag, tmp, num = make(map[int32]struct{}), int32(0), int32(0)
	for i := 0; i < k && i < len(s); i++ {
		num, tmp = num<<1+int32(s[i]-'0'), tmp<<1+1
	}
	flag[num] = struct{}{}

	for i := k; i < len(s); i++ {
		num = num<<1&tmp + int32(s[i]-'0')
		flag[num] = struct{}{}
	}
	return len(flag) == 1<<k
}
