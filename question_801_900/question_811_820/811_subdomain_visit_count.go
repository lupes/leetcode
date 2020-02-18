package question_811_820

import (
	"strconv"
)

// 811. 子域名访问计数
// https://leetcode-cn.com/problems/subdomain-visit-count
// Topics: 哈希表

func subdomainVisits(cpdomains []string) []string {
	var domains = make(map[string]int)
	for _, domain := range cpdomains {
		num := 0
		for i, c := range domain {
			if c >= '0' && c <= '9' {
				num = num*10 + int(c-'0')
			}
			if c == ' ' || c == '.' {
				domains[domain[i+1:]] += num
			}
		}
	}
	var res = make([]string, 0, len(domains))
	for key, v := range domains {
		res = append(res, strconv.Itoa(v)+" "+key)
	}
	return res
}
