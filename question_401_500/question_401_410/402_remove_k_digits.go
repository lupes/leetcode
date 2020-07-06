package question_401_410

// 402. 移掉K位数字
// https://leetcode-cn.com/problems/remove-k-digits
// Topics: 栈 贪心算法

func removeKdigits(num string, k int) string {
	if len(num) == k {
		return "0"
	}
	var stack, res = make([]uint8, 0, len(num)), ""
	for i := 0; i < len(num); {
		if len(stack) > 0 && k > 0 && stack[len(stack)-1] > num[i] {
			stack = stack[:len(stack)-1]
			k--
		} else {
			if len(stack) != 0 || num[i] != '0' {
				stack = append(stack, num[i])
			}
			i++
		}
	}
	res = string(stack)
	if k > 0 && k < len(res) {
		res = res[:len(res)-k]
	} else if len(res) == 0 {
		res = "0"
	}
	return res
}
