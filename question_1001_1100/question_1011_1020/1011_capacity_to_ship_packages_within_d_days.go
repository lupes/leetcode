package question_1011_1020

// 1011. 在 D 天内送达包裹的能力
// https://leetcode-cn.com/problems/capacity-to-ship-packages-within-d-days
// Topics: 数组 二分查找

func shipWithinDays(weights []int, D int) int {
	min, max := 0, 0
	for _, w := range weights {
		if w > min {
			min = w
		}
		max += w
	}
	for max > min {
		weight := (max + min) / 2
		d := shipWithinDaysHelper(weights, weight)

		if d > D {
			min = weight + 1
		} else {
			max = weight
		}
	}
	return min
}

func shipWithinDaysHelper(weights []int, weight int) int {
	var r, t = 0, 0
	for i := 0; i < len(weights); {
		if t+weights[i] <= weight {
			t += weights[i]
			i++
		} else {
			t = 0
			r++
		}
	}
	if t > 0 {
		r++
	}
	return r
}
