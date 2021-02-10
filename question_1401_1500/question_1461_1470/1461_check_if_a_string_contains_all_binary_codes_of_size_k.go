package question_1461_1470

// 1451. 重新排列句子中的单词
// https://leetcode-cn.com/problems/rearrange-words-in-a-sentence/
// Topics: 排序 字符串

func hasAllCodes(s string, k int) bool {
	var flag, num = make(map[uint32]struct{}), uint32(0)
	for i := 0; i < k && i < len(s); i++ {
		num = num<<1 + uint32(s[i]-'0')
	}
	flag[num] = struct{}{}

	for i := k; i < len(s); i++ {
		num = (num<<(33-k))>>(32-k) + uint32(s[i]-'0')
		flag[num] = struct{}{}
	}
	return len(flag) == 1<<k
}
