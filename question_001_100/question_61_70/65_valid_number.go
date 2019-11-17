package question_61_70

import (
	"fmt"
	"strings"
)

// 65. 有效数字
// https://leetcode-cn.com/problems/valid-number/

func isNumber(s string) bool {
	s = strings.Trim(s, " ")
	if len(s) == 0 {
		fmt.Println("长度为0")
		return false
	}
	for _, r := range s {
		if !(r >= '0' && r <= '9') && r != '+' && r != '-' && r != 'e' && r != '.' {
			fmt.Println("有非法字符")
			return false
		}
	}
	arr := strings.Split(s, "e")
	if len(arr) > 2 {
		fmt.Println("e出现次数多余一次")
		return false
	}
	for _, a := range arr {
		if len(a) == 0 {
			fmt.Println("e左右两侧不能没有数字")
			return false
		}
	}
	if len(arr) == 2 {
		t := arr[1]
		if len(t) == 1 && (t == "." || t == "+" || t == "-") {
			fmt.Println("e右侧长度为1且是符号")
			return false
		}
		if strings.Contains(t, ".") {
			fmt.Println("e右侧不能是小数")
			return false
		}
		if i := strings.Count(t, "+"); i > 1 {
			fmt.Println("e右侧+出现次数多余一次")
			return false
		}
		if i := strings.Index(t, "+"); i != -1 && i != 0 {
			fmt.Println("e右侧+出现位置必须是第一位")
			return false
		}
		if i := strings.Count(t, "-"); i > 1 {
			fmt.Println("e右侧-出现次数多余一次")
			return false
		}
		if i := strings.Index(t, "-"); i != -1 && i != 0 {
			fmt.Println("e右侧-出现位置必须是第一位")
			return false
		}
	}
	if len(arr) > 0 {
		t := arr[0]
		if len(t) == 1 && (t == "." || t == "+" || t == "-") {
			fmt.Println("e左侧长度为1且是符号")
			return false
		}
		if len(t) == 2 && (t == "-." || t == "+.") {
			fmt.Println("e右侧长度为2且是符号")
			return false
		}
		if i := strings.Count(t, "."); i > 1 {
			fmt.Println("e左侧.出现次数多余一次")
			return false
		}
		if i := strings.Count(t, "+"); i > 1 {
			fmt.Printf("e左侧+出现次数多余一次:%d\n", i)
			return false
		}
		if i := strings.Index(t, "+"); i != -1 && i != 0 {
			fmt.Printf("e左侧+出现位置必须是第一位:%d\n", i)
			return false
		}
		if i := strings.Count(t, "-"); i > 1 {
			fmt.Println("e左侧-出现次数多余一次")
			return false
		}
		if i := strings.Index(t, "-"); i != -1 && i != 0 {
			fmt.Println("e左侧-出现位置必须是第一位")
			return false
		}
	}
	return true
}
