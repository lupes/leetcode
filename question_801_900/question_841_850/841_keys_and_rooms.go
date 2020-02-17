package question_841_850

// 841. 钥匙和房间
// https://leetcode-cn.com/problems/keys-and-rooms
// Topics: 深度优先搜索 图

func canVisitAllRooms(rooms [][]int) bool {
	var flag = map[int]struct{}{0: {}}
	var next = []int{0}
	for len(next) > 0 {
		keys := rooms[next[0]]
		for _, key := range keys {
			if _, ok := flag[key]; !ok {
				flag[key] = struct{}{}
				next = append(next, key)
			}
		}
		next = next[1:]
	}
	return len(flag) == len(rooms)
}
