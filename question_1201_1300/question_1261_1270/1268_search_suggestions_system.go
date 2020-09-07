package question_1261_1270

import (
	"sort"
)

// 1262. 可被三整除的最大和
// https://leetcode-cn.com/problems/greatest-sum-divisible-by-three/
// Topics: 动态规划

func suggestedProducts(products []string, searchWord string) [][]string {
	var res = make([][]string, len(searchWord))
	sort.Strings(products)
	for i := range searchWord {
		for _, product := range products {
			if len(product) > i && searchWord[:i+1] == product[:i+1] {
				res[i] = append(res[i], product)
				if len(res[i]) == 3 {
					break
				}
			}
		}
		if len(res[i]) == 0 {
			return res
		}
	}
	return res
}
