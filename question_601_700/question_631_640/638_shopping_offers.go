package question_631_640

// 638. 大礼包
// https://leetcode-cn.com/problems/shopping-offers
// Topics: 深度优先搜索 动态规划

func shoppingOffers(price []int, special [][]int, needs []int) int {
	var n = len(price)
	var filterSpecial [][]int
	for _, tmp := range special {
		if shoppingOffersSum(price, tmp) > tmp[n] {
			filterSpecial = append(filterSpecial, tmp)
		}
	}

	return shoppingOffersDfs(price, filterSpecial, needs, n)
}

func shoppingOffersValid(special []int, needs []int) ([]int, bool) {
	var res = make([]int, len(needs))
	for i := range needs {
		if needs[i] < special[i] {
			return nil, false
		}
		res[i] = needs[i] - special[i]
	}
	return res, true
}

func shoppingOffersSum(price []int, needs []int) int {
	var sum = 0
	for i := range price {
		sum += needs[i] * price[i]
	}
	return sum
}

func shoppingOffersDfs(price []int, special [][]int, needs []int, n int) int {
	var min = shoppingOffersSum(price, needs)
	for _, now := range special {
		if newNeeds, ok := shoppingOffersValid(now, needs); ok {
			tmp := now[n] + shoppingOffersDfs(price, special, newNeeds, n)
			if tmp < min {
				min = tmp
			}
		}
	}
	return min
}
