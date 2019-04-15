package question_51_60

import "strings"

func lengthOfLastWord(s string) int {
	if s == "" {
		return 0
	}
	s = strings.TrimRight(s, " ")
	size := len(s)
	for i := size - 1; i >= 0; i-- {
		if s[i] == ' ' {
			return len(s[i+1:])
		}
	}
	return size
}
