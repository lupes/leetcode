package question_931_940

// 939. 最小面积矩形
// https://leetcode-cn.com/problems/minimum-area-rectangle
// Topics: 哈希表

func minAreaRect(points [][]int) int {
	var flag = make(map[int]map[int]struct{})
	for _, point := range points {
		if _, ok := flag[point[0]]; !ok {
			flag[point[0]] = make(map[int]struct{})
		}
		flag[point[0]][point[1]] = struct{}{}
	}
	var min, f = 0, false
	for i := 0; i < len(points)-1; i++ {
		for j := i + 1; j < len(points); j++ {
			if points[i][0] != points[j][0] && points[i][1] != points[j][1] {
				_, ok1 := flag[points[i][0]][points[j][1]]
				_, ok2 := flag[points[j][0]][points[i][1]]
				if ok1 && ok2 {
					tmp := diff(points[i][0], points[j][0]) * diff(points[i][1], points[j][1])
					if !f || tmp < min {
						f = true
						min = tmp
					}
				}
			}
		}
	}
	return min
}

func diff(x, y int) int {
	if x > y {
		return x - y
	} else {
		return y - x
	}
}
