package question_11_20

func maxArea(height []int) int {
	size := len(height)
	var h, max, tmp, j int
	for i := 0; i < size-1; i++ {
		for j = i + 1; j < size; j++ {
			h = height[i]
			if height[i] > height[j] {
				h = height[j]
			}
			tmp = h * (j - i)
			if tmp > max {
				max = tmp
			}
		}
	}
	return max
}
