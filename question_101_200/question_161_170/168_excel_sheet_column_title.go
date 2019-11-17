package question_161_170

// 168. Excel表列名称
// https://leetcode-cn.com/problems/excel-sheet-column-title

func convertToTitle(n int) string {
	res := ""
	for n > 0 {
		t := (n - 1) % 26
		n = n / 26
		res = string(t+'A') + res
		if t == 25 {
			n--
		}
	}
	return res
}

func ConvertToTitle(n int) string {
	return convertToTitle(n)
}
