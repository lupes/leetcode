package question_61_70

func plusOne(digits []int) []int {
	size, height, tmp := len(digits), 1, 0
	for i := size - 1; i >= 0; i-- {
		tmp = digits[i] + height
		digits[i] = tmp % 10
		height = tmp / 10
		if height == 0 {
			break
		}
	}
	if height == 1 {
		digits = append([]int{1}, digits...)
	}
	return digits
}
