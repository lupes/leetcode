package question_1291_1300

// 1299. 将每个元素替换为右侧最大元素
// https://leetcode-cn.com/problems/replace-elements-with-greatest-element-on-right-side/
// Topics: 数组

func replaceElements(arr []int) []int {
	var max, n = -1, 0
	for i := len(arr) - 1; i >= 0; i-- {
		arr[i], n = max, arr[i]
		if n > max {
			max = n
		}
	}
	return arr
}
