package question_1081_1090

// 1089. 复写零
// https://leetcode-cn.com/problems/duplicate-zeros
// Topics: 数组

func duplicateZeros(arr []int) {
	var t, l, index = 0, len(arr), 0
	for i := 0; i < len(arr); i++ {
		if arr[i] == 0 {
			t++
		}
		t++
		if t >= l {
			index = i
			break
		}
	}
	if index == l-1 {
		return
	}
	if t > l {
		arr[l-1] = arr[index]
		l--
		index--
	}
	for i, j := index, l-1; i >= 0; i, j = i-1, j-1 {
		arr[j] = arr[i]
		if arr[i] == 0 {
			j--
			arr[j] = arr[i]
		}
	}
}
