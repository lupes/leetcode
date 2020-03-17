package question_921_930

// 925. 长按键入
// https://leetcode-cn.com/problems/long-pressed-name
// Topics: 双指针 字符串

func isLongPressedName(name string, typed string) bool {
	if name == typed {
		return true
	}
	var NFlag, TFlag = make([][2]int32, 0), make([][2]int32, 0)
	for _, c := range name {
		l := len(NFlag)
		if l > 0 && c == NFlag[l-1][0] {
			NFlag[l-1][1]++
		} else {
			NFlag = append(NFlag, [2]int32{c, 1})
		}
	}
	for _, c := range typed {
		l := len(TFlag)
		if l > 0 && c == TFlag[l-1][0] {
			TFlag[l-1][1]++
		} else {
			TFlag = append(TFlag, [2]int32{c, 1})
		}
	}
	if len(NFlag) != len(TFlag) {
		return false
	}
	for i, flag := range NFlag {
		t := TFlag[i]
		if t[0] != flag[0] || t[1] < flag[1] {
			return false
		}
	}
	return true
}
