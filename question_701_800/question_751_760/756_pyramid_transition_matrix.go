package question_751_760

// 756. 金字塔转换矩阵
// https://leetcode-cn.com/problems/pyramid-transition-matrix
// Topics: 位运算 深度优先搜索

func pyramidTransition(bottom string, allowed []string) bool {
	var flag = make(map[string]struct{})
	for _, allow := range allowed {
		flag[allow] = struct{}{}
	}
	return pyramidTransitionHelper(bottom, "", flag, 0)
}

func pyramidTransitionHelper(bottom, newBottom string, allowed map[string]struct{}, i int) bool {
	if len(bottom) == 1 {
		return true
	}
	if len(newBottom) == len(bottom)-1 {
		return pyramidTransitionHelper(newBottom, "", allowed, 0)
	}
	for j := 'A'; j < 'H'; j++ {
		if _, ok := allowed[bottom[i:i+2]+string(j)]; ok {
			if pyramidTransitionHelper(bottom, newBottom+string(j), allowed, i+1) {
				return true
			}
		}
	}
	return false
}
