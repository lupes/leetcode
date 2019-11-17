package question_11_20

// 11. 盛最多水的容器
// https://leetcode-cn.com/problems/container-with-most-water
// Topics: 数组 双指针

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
