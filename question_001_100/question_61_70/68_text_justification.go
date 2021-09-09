package question_61_70

import (
	"strings"
)

// 68. 文本左右对齐
// https://leetcode-cn.com/problems/text-justification
// Topics: 字符串 模拟

func fullJustify(words []string, maxWidth int) []string {
	var lines []string
	var l = 0
	for left, right := 0, 0; right < len(words); {
		if right == len(words)-1 && l+len(words[right]) < maxWidth {
			s := strings.Join(words[left:right+1], " ")
			s += strings.Repeat(" ", maxWidth-len(s))
			lines = append(lines, s)
			break
		} else if l+len(words[right]) < maxWidth {
			l += len(words[right]) + 1
			right++
		} else if l+len(words[right]) == maxWidth {
			lines = append(lines, strings.Join(words[left:right+1], " "))
			l = 0
			right++
			left = right
		} else {
			s := align(words[left:right], maxWidth)
			lines = append(lines, s)
			l = 0
			left = right
		}
	}

	return lines
}

func align(words []string, width int) string {
	l := len(words)
	if l == 1 {
		return words[0] + strings.Repeat(" ", width-len(words[0]))
	}
	l1 := len(strings.Join(words, ""))
	re := (width - l1) % (l - 1)
	qu := (width - l1) / (l - 1)
	s := ""
	for i, word := range words {
		if i == len(words)-1 {
			s += word
		} else if i < re {
			s += word + strings.Repeat(" ", qu+1)
		} else {
			s += word + strings.Repeat(" ", qu)
		}
	}
	return s
}
