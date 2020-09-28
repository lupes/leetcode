package question_381_390

// 381. O(1) 时间插入、删除和获取随机元素 - 允许重复
// https://leetcode-cn.com/problems/insert-delete-getrandom-o1-duplicates-allowed
// Topics: 设计 数组 哈希表

type RandomizedCollection struct {
}

/** Initialize your data structure here. */
func Constructor1() RandomizedCollection {
	return RandomizedCollection{}
}

/** Inserts a value to the collection. Returns true if the collection did not already contain the specified element. */
func (this *RandomizedCollection) Insert(val int) bool {
	return false
}

/** Removes a value from the collection. Returns true if the collection contained the specified element. */
func (this *RandomizedCollection) Remove(val int) bool {
	return false
}

/** Get a random element from the collection. */
func (this *RandomizedCollection) GetRandom() int {
	return 0
}

/**
 * Your RandomizedCollection object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Insert(val);
 * param_2 := obj.Remove(val);
 * param_3 := obj.GetRandom();
 */
