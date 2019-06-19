package question_121_130

func maxProfit(prices []int) int {
	max := 0
	for i, p := range prices {
		for _, t := range prices[i+1:] {
			s := t - p
			if s > max {
				max = s
			}
		}
	}
	return max
}
