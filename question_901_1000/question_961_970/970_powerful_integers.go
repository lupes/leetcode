package question_961_970

// 970. 强整数
// https://leetcode-cn.com/problems/powerful-integers
// Topics: 哈希表 数学

func powerfulIntegers(x int, y int, bound int) []int {
	var flagx = map[int]struct{}{1: {}}
	var flagy = map[int]struct{}{1: {}}
	var flagz = map[int]struct{}{}
	var tmp = x
	for tmp < bound && x != 1 {
		flagx[tmp] = struct{}{}
		tmp *= x
	}
	tmp = y
	for tmp < bound && y != 1 {
		flagy[tmp] = struct{}{}
		tmp *= y
	}
	var res []int
	for i, _ := range flagx {
		for j, _ := range flagy {
			if i+j <= bound {
				flagz[i+j] = struct{}{}
			}
		}
	}
	for i, _ := range flagz {
		res = append(res, i)
	}
	return res
}
