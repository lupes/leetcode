package question_701_710

// 706. 设计哈希映射
// https://leetcode-cn.com/problems/design-hashmap
// Topics: 设计 哈希表

const max = 1e6

type MyHashMap struct {
	max int
	Map [][][]int
}

/** Initialize your data structure here. */
func Constructor4() MyHashMap {
	return MyHashMap{
		max: -1,
		Map: make([][][]int, 100),
	}
}

func (this *MyHashMap) index(key int) (int, int, int) {
	return key / 10000, (key % 10000) / 100, key % 100
}

/** value will always be non-negative. */
func (this *MyHashMap) Put(key int, value int) {
	if key == max {
		this.max = value
	} else {
		a, b, c := this.index(key)
		if this.Map[a] == nil {
			this.Map[a] = make([][]int, 100)
		}
		if this.Map[a][b] == nil {
			this.Map[a][b] = make([]int, 100)
			for i := range this.Map[a][b] {
				this.Map[a][b][i] = -1
			}
		}
		this.Map[a][b][c] = value
	}
}

/** Returns the value to which the specified key is mapped, or -1 if this map contains no mapping for the key */
func (this *MyHashMap) Get(key int) int {
	if key == max {
		return this.max
	}
	a, b, c := this.index(key)
	if this.Map[a] == nil {
		return -1
	}
	if this.Map[a][b] == nil {
		return -1
	}
	return this.Map[a][b][c]
}

/** Removes the mapping of the specified value key if this map contains a mapping for the key */
func (this *MyHashMap) Remove(key int) {
	if key == max {
		this.max = -1
		return
	}
	a, b, c := this.index(key)
	if this.Map[a] == nil {
		return
	}
	if this.Map[a][b] == nil {
		return
	}
	this.Map[a][b][c] = -1
}
