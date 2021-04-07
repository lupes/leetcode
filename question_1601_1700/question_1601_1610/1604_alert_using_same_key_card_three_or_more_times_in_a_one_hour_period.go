package question_1601_1610

import (
	"sort"
)

// 1604. 警告一小时内使用相同员工卡大于等于三次的人
// https://leetcode-cn.com/problems/alert-using-same-key-card-three-or-more-times-in-a-one-hour-period/
// Topics: 字符串 OrderedMap

func alertNames(keyName []string, keyTime []string) []string {
	var flag = make(map[string][]int)
	for i, t := range keyTime {
		flag[keyName[i]] = append(flag[keyName[i]], (int(t[0]-'0')*10+int(t[1]-'0'))*60+int(t[3]-'0')*10+int(t[4]-'0'))
	}

	var names []string
	for name, times := range flag {
		sort.Ints(times)
		for i := 2; i < len(times); i++ {
			if times[i] <= times[i-2]+60 {
				names = append(names, name)
				goto Next
			}
		}
	Next:
	}
	sort.Strings(names)
	return names
}
