package question_601_610

// 605. 种花问题
// https://leetcode-cn.com/problems/can-place-flowers/
// Topics: 数组

func canPlaceFlowers(flowerbed []int, n int) bool {
	m := 1
	for _, t := range flowerbed {
		if t == 1 {
			n -= (m - 1) / 2
			m = 0
		} else {
			m++
		}
	}
	n -= m / 2
	return n < 1
}
