package question_1261_1270

// 1276. 不浪费原料的汉堡制作方案
// https://leetcode-cn.com/problems/number-of-burgers-with-no-waste-of-ingredients/
// Topics: 贪心算法 数学

func numOfBurgers(tomatoSlices int, cheeseSlices int) []int {
	if tomatoSlices%2 != 0 || cheeseSlices*4 < tomatoSlices || cheeseSlices*2 > tomatoSlices {
		return nil
	}

	a := (tomatoSlices - cheeseSlices*2) / 2
	b := cheeseSlices - a
	return []int{a, b}
}
