package question_371_380

import (
	"math/rand"
)

// 380. 常数时间插入、删除和获取随机元素
// https://leetcode-cn.com/problems/insert-delete-getrandom-o1
// Topics: 设计 数组 哈希表

type RandomizedSet struct {
	MapData map[int]int
	ArrData []int
	Len     int
}

/** Initialize your data structure here. */
func Constructor() RandomizedSet {
	return RandomizedSet{
		MapData: make(map[int]int),
		ArrData: nil,
		Len:     0,
	}
}

/** Inserts a value to the set. Returns true if the set did not already contain the specified element. */
func (this *RandomizedSet) Insert(val int) bool {
	if _, ok := this.MapData[val]; ok {
		return false
	}
	this.MapData[val] = len(this.ArrData)
	this.ArrData = append(this.ArrData, val)
	this.Len += 1
	return true
}

/** Removes a value from the set. Returns true if the set contained the specified element. */
func (this *RandomizedSet) Remove(val int) bool {
	index, ok := this.MapData[val]
	if !ok {
		return false
	}
	this.ArrData[index] = this.ArrData[this.Len-1]
	this.MapData[this.ArrData[index]] = index
	this.Len -= 1
	this.ArrData = this.ArrData[:this.Len]
	delete(this.MapData, val)
	return true
}

/** Get a random element from the set. */
func (this *RandomizedSet) GetRandom() int {
	n := rand.Intn(this.Len)
	return this.ArrData[n]
}
