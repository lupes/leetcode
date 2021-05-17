package question_1691_1700

// 1694. 重新格式化电话号码
// https://leetcode-cn.com/problems/reformat-phone-number/
// Topics: 字符串

func reformatNumber(number string) string {
	var res, tmp = "", make([]byte, 0, 3)
	for i, c := range number {
		if c >= '0' && c <= '9' {
			tmp = append(tmp, number[i])
			if len(tmp) > 2 {
				res += "-" + string(tmp)
				tmp = make([]byte, 0, 3)
			}
		}
	}
	if len(tmp) == 2 {
		res += "-" + string(tmp)
	} else if len(tmp) == 1 {
		t := res[len(res)-3:] + string(tmp)
		res = res[:len(res)-3] + t[:2] + "-" + t[2:]
	}
	return res[1:]
}
