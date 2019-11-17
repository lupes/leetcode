package question_161_170

import (
	"strconv"
	"strings"
)

// 165. 比较版本号
// https://leetcode-cn.com/problems/compare-version-numbers
// Topics: 字符串

func compareVersion(version1 string, version2 string) int {
	vers1 := strings.Split(version1, ".")
	vers2 := strings.Split(version2, ".")
	size1, size2 := len(vers1), len(vers2)
	maxSize := size1
	if size2 > size1 {
		maxSize = size2
	}
	for i := 0; i < maxSize; i++ {
		val1, val2 := 0, 0
		if i < size1 {
			val1, _ = strconv.Atoi(vers1[i])
		}
		if i < size2 {
			val2, _ = strconv.Atoi(vers2[i])
		}
		if val1 > val2 {
			return 1
		} else if val1 < val2 {
			return -1
		}
	}
	return 0
}
