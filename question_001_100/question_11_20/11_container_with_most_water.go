package question_11_20

// 11. 盛最多水的容器
// https://leetcode-cn.com/problems/container-with-most-water
// Topics: 数组 双指针

func maxArea(height []int) int {
	max, h := 0, 0
	for left, right := 0, len(height)-1; left < right; {
		w := right - left
		if height[left] > height[right] {
			h = height[right]
			right--
		} else {
			h = height[left]
			left++
		}
		if h*w > max {
			max = h * w
		}
	}
	return max
}
