package question_1311_1320

import (
	"sort"
)

// 1333. 餐厅过滤器
// https://leetcode-cn.com/problems/filter-restaurants-by-vegan-friendly-price-and-distance/
// Topics: 排序 数组

func filterRestaurants(restaurants [][]int, veganFriendly int, maxPrice int, maxDistance int) []int {
	var tmp [][]int
	for _, restaurant := range restaurants {
		if veganFriendly == 1 && restaurant[2] == 0 {
			continue
		}
		if restaurant[3] > maxPrice {
			continue
		}
		if restaurant[4] > maxDistance {
			continue
		}
		tmp = append(tmp, restaurant)
	}

	sort.Slice(tmp, func(i, j int) bool {
		return tmp[i][1] > tmp[j][1] || (tmp[i][1] == tmp[j][1] && tmp[i][0] > tmp[j][0])
	})

	var res []int
	for _, item := range tmp {
		res = append(res, item[0])
	}
	return res
}
