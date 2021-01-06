package question_1301_1310

// 1306. 跳跃游戏 III
// https://leetcode-cn.com/problems/all-elements-in-two-binary-search-trees/
// Topics: 排序 树

func canReach(arr []int, start int) bool {
	if start < 0 || start >= len(arr) {
		return false
	}
	if arr[start] == -1 {
		return false
	}
	if arr[start] == 0 {
		return true
	}
	var tmp, res = arr[start], false
	arr[start] = -1
	res = res || canReach(arr, start+tmp)
	res = res || canReach(arr, start-tmp)
	return res
}
