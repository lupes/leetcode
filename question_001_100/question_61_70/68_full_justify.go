package question_61_70

import (
	"strings"
)

// 68. 文本左右对齐
// https://leetcode-cn.com/problems/text-justification

func fullJustify(words []string, maxWidth int) []string {
	var line []string
	size := len(words)
	for start := 0; start <= size; {
		last := start
		for ; last < size; last++ {
			if len(strings.Join(words[start:last+1], " ")) > maxWidth {
				break
			}
		}
		if last < size {
			s := align(words[start:last], maxWidth)
			line = append(line, s)
		} else {
			s := align(words[start:last], maxWidth, true)
			line = append(line, s)
			break
		}
		start = last
	}
	//for i := 0; i < size; i++ {
	//	tmp = append(tmp, words[i])
	//	s := strings.Join(tmp, " ")
	//	if len(s) > maxWidth {
	//		tmp = tmp[:len(tmp)-1]
	//		if i < size-1 {
	//			s := align(tmp, maxWidth)
	//			line = append(line, s)
	//		} else {
	//			s := align(tmp, maxWidth)
	//			line = append(line, s)
	//			s = align([]string{words[i]}, maxWidth, true)
	//			line = append(line, s)
	//		}
	//		tmp = []string{words[i]}
	//	} else if i == size-1 {
	//		s := align(tmp, maxWidth, true)
	//		line = append(line, s)
	//	}
	//}
	return line
}

func align(words []string, width int, last ...bool) string {
	if len(last) == 1 && last[0] {
		words = []string{strings.Join(words, " ")}
	}
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
			break
		}
		if i < re {
			s += word + strings.Repeat(" ", qu+1)
		} else {
			s += word + strings.Repeat(" ", qu)
		}
	}
	return s
}
