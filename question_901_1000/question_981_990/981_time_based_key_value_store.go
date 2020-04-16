package question_981_990

// 981. 基于时间的键值存储
// https://leetcode-cn.com/problems/time-based-key-value-store
// Topics: 哈希表 二分查找

type Value struct {
	timestamp int
	value     string
}

type TimeMap struct {
	m map[string][]Value
}

/** Initialize your data structure here. */
func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]Value),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	this.m[key] = append(this.m[key], Value{
		timestamp: timestamp,
		value:     value,
	})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	m, ok := this.m[key]
	if !ok {
		return ""
	}
	if m[0].timestamp > timestamp {
		return ""
	}
	left, right, center := 0, len(m), 0
	for right > left {
		center = (left + right) / 2
		if m[center].timestamp == timestamp {
			break
		} else if m[center].timestamp > timestamp {
			right = center
		} else {
			left = center + 1
		}
	}
	return m[center].value
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */
