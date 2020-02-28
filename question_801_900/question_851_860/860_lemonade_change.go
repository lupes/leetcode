package question_851_860

// 860. 柠檬水找零
// https://leetcode-cn.com/problems/lemonade-change
// Topics: 贪心算法

func lemonadeChange(bills []int) bool {
	var a, b = 0, 0
	for _, bill := range bills {
		switch {
		case bill == 5:
			a++
		case bill == 10 && a > 0:
			a, b = a-1, b+1
		case bill == 20 && b > 0 && a > 0:
			a, b = a-1, b-1
		case bill == 20 && a > 2:
			a -= 3
		default:
			return false
		}
	}
	return true
}
