package question_691_700

import (
	"sort"
)

// 692. 前K个高频单词
// https://leetcode-cn.com/problems/top-k-frequent-words
// Topics: 堆 字典树 哈希表

func topKFrequent(words []string, k int) []string {
	var max = 0
	var flag1 = make(map[string]int)
	var flag2 = make(map[int][]string)
	for _, word := range words {
		flag1[word] += 1
		if flag1[word] > max {
			max = flag1[word]
		}
	}
	for k, v := range flag1 {
		flag2[v] = append(flag2[v], k)
	}
	var res []string
	for k > 0 {
		v, ok := flag2[max]
		if ok {
			sort.Strings(v)
			l := len(v)
			if k >= l {
				res = append(res, v...)
				k -= l
			} else {
				res = append(res, v[:k]...)
				k -= k
			}
		}
		max--
	}
	return res
}
