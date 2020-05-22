package question_1071_1080

// 1079. 活字印刷
// https://leetcode-cn.com/problems/letter-tile-possibilities
// Topics: 回溯算法

func numTilePossibilities(tiles string) int {
	var flag = make([]bool, len(tiles))
	var m = make(map[string]struct{})
	numTilePossibilitiesHelp(tiles, "", len(tiles), flag, m)
	return len(m) - 1
}

func numTilePossibilitiesHelp(titles, prefix string, l int, flag []bool, m map[string]struct{}) {
	if l == 0 {
		m[prefix] = struct{}{}
		return
	}
	for i := 0; i < len(titles); i++ {
		if flag[i] == false {
			flag[i] = true
			numTilePossibilitiesHelp(titles, prefix, l-1, flag, m)
			numTilePossibilitiesHelp(titles, prefix+titles[i:i+1], l-1, flag, m)
			flag[i] = false
		}
	}
}
