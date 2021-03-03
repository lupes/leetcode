package question_1491_1500

// 1496. 判断路径是否相交
// https://leetcode-cn.com/problems/path-crossing/
// Topics: 字符串

func isPathCrossing(path string) bool {
	var flag = make(map[int32]int, 4)
	for i, c := range path {
		if i > 0 && (path[i-1:i+1] == "NS" || path[i-1:i+1] == "SN" || path[i-1:i+1] == "EW" || path[i-1:i+1] == "WE") {
			return true
		}
		flag[c]++
		if flag['N'] == flag['S'] && flag['E'] == flag['W'] {
			return true
		}
	}
	return false
}
