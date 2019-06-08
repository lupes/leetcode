package question_91_100

import (
	"strconv"
)

func restoreIpAddresses(s string) []string {
	var res []string
	var l = len(s)
	for i := 1; i < l-2 && i < 4; i++ {
		for j := i + 1; j < l-1 && j < i+4; j++ {
			for k := j + 1; k < l && k < j+4; k++ {
				if len(s[k:]) > 3 {
					continue
				}
				if ip, ok := validIpAddresses(s, i, j, k); ok {
					res = append(res, ip)
				}
			}
		}
	}
	return res
}

func validIpAddresses(s string, i, j, k int) (string, bool) {
	ip := []string{s[0:i], s[i:j], s[j:k], s[k:]}
	for _, t := range ip {
		if len(t) > 3 || (len(t) != 1 && t[0] == '0') {
			return "", false
		}
		res, _ := strconv.ParseInt(t, 10, 32)
		if res > 255 {
			return "", false
		}
	}
	return ip[0] + "." + ip[1] + "." + ip[2] + "." + ip[3], true
}
