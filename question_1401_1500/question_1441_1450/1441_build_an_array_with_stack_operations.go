package question_1441_1450

// 1441. 用栈操作构建数组
// https://leetcode-cn.com/problems/build-an-array-with-stack-operations/
// Topics: 栈

func buildArray(target []int, n int) []string {
	var res []string
	var left = 0
	for i := 1; i <= n && left < len(target); i++ {
		if target[left] == i {
			res = append(res, "Push")
			left++
		} else {
			res = append(res, "Push", "Pop")
		}
	}
	return res
}
