package question_41_50

func trap(height []int) int {
	size := len(height)
	var l, r, res, h, max int
	for _, t := range height {
		if t > max {
			max = t
		}
	}
	for {
		l, r = 0, size-1
		h += 1
		if max < h {
			break
		}
		for l < r {
			if height[l] < h {
				l++
			}
			if height[r] < h {
				r--
			}
			if height[l] >= h && height[r] >= h {
				break
			}
		}
		for i := l; i < r; i++ {
			if height[i] < h {
				res += 1
			}
		}
	}
	return res
}
