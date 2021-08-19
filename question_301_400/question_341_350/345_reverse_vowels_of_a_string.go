package question_341_350

// 345. 反转字符串中的元音字母
// https://leetcode-cn.com/problems/reverse-vowels-of-a-string/
// Topics: 双指针 字符串

var flag = map[byte]struct{}{
	'a': {}, 'o': {}, 'e': {}, 'i': {}, 'u': {}, 'A': {}, 'O': {}, 'E': {}, 'I': {}, 'U': {},
}

func reverseVowels(s string) string {
	bs := []byte(s)
	left, right := 0, len(bs)-1
	for right > left {
		_, leftOk := flag[bs[left]]
		_, rightOk := flag[bs[right]]
		if leftOk && rightOk {
			bs[left], bs[right] = bs[right], bs[left]
			left++
			right--
		} else if !leftOk {
			left++
		} else if !rightOk {
			right--
		}
	}
	return string(bs)
}
