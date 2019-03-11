package question_01_10

import "fmt"

func isMatch(s string, p string) bool {
	if s == "" && p == "" {
		return true
	}
	if s != "" && p == "" {
		return false
	}
	sHead, sLen, pHead, pLen := 0, len(s), 0, len(p)
	fmt.Printf("%d %d\n", sHead, sLen)
	var tmp []string
	for pHead < pLen {
		if ('a' <= p[pHead] && p[pHead] <= 'z') || p[pHead] == '.' {
			if pHead+1 < pLen && p[pHead+1] == '*' {
				if len(tmp) == 0 || tmp[len(tmp)-1] != p[pHead:pHead+2] {
					tmp = append(tmp, p[pHead:pHead+2])
				}
				pHead += 2
			} else {
				tmp = append(tmp, p[pHead:pHead+1])
				tLen := len(tmp)
				if tLen >= 2 && tmp[tLen-1][:1] == p[pHead:pHead+1] {
					tmp[tLen-2], tmp[tLen-1] = tmp[tLen-1], tmp[tLen-2]
				}
				pHead += 1
			}
		}
	}
	for _, t := range tmp {
		if len(t) > 1 {
			for {
				if s[sHead] == t[0] && sHead+1 < sLen {
					sHead += 1
				} else {
					break
				}
			}
		} else {
			if t == "." {
				sHead += 1
			} else {
				if s[sHead] == t[0] {
					sHead += 1
				} else {
					return false
				}
			}
		}
	}
	fmt.Printf("%v\n", tmp)
	return false
}
