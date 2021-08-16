package question_381_390

import (
	"math/rand"
)

// 381. O(1) 时间插入、删除和获取随机元素 - 允许重复
// https://leetcode-cn.com/problems/insert-delete-getrandom-o1-duplicates-allowed
// Topics: 设计 数组 哈希表

type RandomizedCollection struct {
	Map  map[int]map[int]struct{}
	Nums []int
	Len  int
}

/** Initialize your data structure here. */
func Constructor1() RandomizedCollection {
	return RandomizedCollection{
		Map:  make(map[int]map[int]struct{}),
		Nums: make([]int, 0, 100),
		Len:  0,
	}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	this.Nums = append(this.Nums, val)
	_, ok := this.Map[val]
	if !ok {
		this.Map[val] = make(map[int]struct{})
	}
	this.Map[val][this.Len] = struct{}{}
	this.Len++
	return !ok
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	_, ok := this.Map[val]
	if ok {
		i := 0
		for i = range this.Map[val] {
			break
		}
		if len(this.Map[val]) == 1 {
			delete(this.Map, val)
		} else {
			delete(this.Map[val], i)
		}
		this.Len--

		if i != this.Len {
			this.Nums[i] = this.Nums[this.Len]
			delete(this.Map[this.Nums[this.Len]], this.Len)
			this.Map[this.Nums[this.Len]][i] = struct{}{}
		}
		this.Nums = this.Nums[:this.Len]
	}
	return ok
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return this.Nums[rand.Intn(this.Len)]
}
