package question_41_50

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	} else {
		return b
	}
}

func trap(height []int) int {
	size := len(height)
	var l, r, res int
	fmt.Printf("%v\n", height)
	for i := 0; i < size; i++ {
		l = max(l, height[i])
		r = max(r, height[size-1-i])
		res += l + r - height[i]
		fmt.Printf("%d %d %d %d\n", i, l, r, res)
	}
	return res - size*l
}
