package question_391_400

// 391. 完美矩形
// https://leetcode-cn.com/problems/perfect-rectangle
// Topics: 数组 扫描线

func isRectangleCover(rectangles [][]int) bool {
	var flag = make(map[[2]int]int)
	var area = 0
	for _, rectangle := range rectangles {
		x1, y1, x2, y2 := rectangle[0], rectangle[1], rectangle[2], rectangle[3]
		flag[[2]int{x1, y1}]++
		flag[[2]int{x1, y2}]++
		flag[[2]int{x2, y1}]++
		flag[[2]int{x2, y2}]++
		area += (y2 - y1) * (x2 - x1)
	}
	x1, y1, x2, y2 := 1000000, 1000000, -1000000, -1000000
	var sum = 0
	for key, value := range flag {
		switch value {
		case 1:
			if key[0] > x2 {
				x2 = key[0]
			} else if key[0] < x1 {
				x1 = key[0]
			}
			if key[1] > y2 {
				y2 = key[1]
			} else if key[1] < y1 {
				y1 = key[1]
			}
			sum++
		case 2, 4:
		default:
			return false
		}
	}
	return (y2-y1)*(x2-x1) == area && sum == 4
}
