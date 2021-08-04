package question_871_880

// 875. 爱吃香蕉的珂珂
// https://leetcode-cn.com/problems/koko-eating-bananas
// Topics: 二分查找

func minEatingSpeed(piles []int, H int) int {
	var min, max = 1, 1
	for _, pile := range piles {
		if pile > max {
			max = pile
		}
	}
	for max > min {
		hour, mid := 0, (min+max)/2
		for _, pile := range piles {
			if pile%mid == 0 {
				hour += pile / mid
			} else {
				hour += pile/mid + 1
			}
		}
		if hour > H {
			min = mid + 1
		} else {
			max = mid
		}
	}
	return min
}
