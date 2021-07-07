package question_671_680

// 677. 键值映射
// https://leetcode-cn.com/problems/map-sum-pairs
// Topics: 字典树

type MapSum struct {
	m map[string]int
}

func Constructor() MapSum {
	return MapSum{
		m: make(map[string]int, 40),
	}
}

func (this *MapSum) Insert(key string, val int) {
	this.m[key] = val
}

func (this *MapSum) Sum(prefix string) int {
	var sum, l = 0, len(prefix)
	for k, v := range this.m {
		if len(k) >= l && k[:l] == prefix {
			sum += v
		}
	}
	return sum
}
