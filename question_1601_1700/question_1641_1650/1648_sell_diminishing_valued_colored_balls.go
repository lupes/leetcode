package question_1641_1650

// 1648. 销售价值减少的颜色球
// https://leetcode-cn.com/problems/sell-diminishing-valued-colored-balls/
// Topics: 贪心算法 排序

import (
	"sort"
)

func maxProfit(inventory []int, orders int) int {
	inventory = append(inventory, 0)
	sort.Slice(inventory, func(i, j int) bool {
		return inventory[i] > inventory[j]
	})

	var res = 0
	for i := 1; i < len(inventory) && orders > 0; i++ {
		d := inventory[i-1] - inventory[i]
		if d*i > orders {
			d = orders / i
		}
		res = (res + ((inventory[i-1]+inventory[i-1]-d+1)*d)/2*i) % (1e9 + 7)
		orders -= d * i
		if d < (inventory[i-1] - inventory[i]) {
			res = (res + orders*(inventory[i-1]-d)) % (1e9 + 7)
			return res
		}

	}
	return res
}
