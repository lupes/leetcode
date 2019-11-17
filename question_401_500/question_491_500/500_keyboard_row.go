package question_491_500

// 500. 键盘行
// https://leetcode-cn.com/problems/keyboard-row/

func findWords(words []string) []string {
	var res []string
	for _, word := range words {
		last, now := 0, 0
		for _, char := range word {
			if char <= 'Z' {
				char += 32
			}
			switch char {
			case 'q', 'w', 'e', 'r', 't', 'y', 'u', 'i', 'o', 'p':
				now = 1
			case 'a', 's', 'd', 'f', 'g', 'h', 'j', 'k', 'l':
				now = 2
			case 'z', 'x', 'c', 'v', 'b', 'n', 'm':
				now = 3
			}
			if last != 0 && now != last {
				goto Next
			}
			last = now
		}
		res = append(res, word)
	Next:
	}
	return res
}
