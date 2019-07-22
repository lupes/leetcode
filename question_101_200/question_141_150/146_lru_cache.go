package question_141_150

// 146. LRU缓存机制
// https://leetcode-cn.com/problems/lru-cache/

type LRUCache struct {
	capacity int
	data     map[int]int
	sorts    []int
}

func Constructor(capacity int) LRUCache {
	return LRUCache{
		capacity: capacity,
		data:     make(map[int]int, capacity),
		sorts:    []int{},
	}
}

func toHead(array []int, key int) []int {
	var i, v int
	for i, v = range array {
		if v == key {
			break
		}
	}
	for j := i; j < len(array)-1; j++ {
		array[j] = array[j+1]
	}
	array[len(array)-1] = key
	return array
}

func (this *LRUCache) Get(key int) int {
	if value, ok := this.data[key]; ok {
		this.sorts = toHead(this.sorts, key)
		return value
	} else {
		return -1
	}
}

func (this *LRUCache) Put(key int, value int) {
	if _, ok := this.data[key]; ok {
		this.sorts = toHead(this.sorts, key)
	} else {
		if len(this.sorts) == this.capacity {
			delete(this.data, this.sorts[0])
			this.sorts = append(this.sorts[1:], key)
		} else {
			this.sorts = append(this.sorts, key)
		}
	}
	this.data[key] = value
}
