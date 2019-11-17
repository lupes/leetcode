package question_341_350

// 345. 反转字符串中的元音字母
// https://leetcode-cn.com/problems/reverse-vowels-of-a-string/
// Topics: 双指针 字符串

func reverseVowels(s string) string {
	var flag = map[byte]struct{}{
		'a': {}, 'o': {}, 'e': {}, 'i': {}, 'u': {}, 'A': {}, 'O': {}, 'E': {}, 'I': {}, 'U': {},
	}
	bs := []byte(s)
	for i, j := 0, len(s)-1; i < j; {
		for ; i < j; i++ {
			if _, ok := flag[bs[i]]; ok {
				break
			}
		}
		for ; j > i; j-- {
			if _, ok := flag[bs[j]]; ok {
				break
			}
		}
		if j > i {
			bs[i], bs[j] = bs[j], bs[i]
		}
		i++
		j--
	}
	return string(bs)
}
