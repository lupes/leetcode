package question_341_350

// 341. 扁平化嵌套列表迭代器
// https://leetcode-cn.com/problems/flatten-nested-list-iterator
// Topics: 栈 设计

type NestedInteger struct {
	flag  bool
	value int
	lists []*NestedInteger
}

func (this NestedInteger) IsInteger() bool {
	return this.flag
}

func (this NestedInteger) GetInteger() int {
	return this.value
}

func (this *NestedInteger) SetInteger(value int) {
	this.value = value
}

func (this *NestedInteger) Add(elem NestedInteger) {
	this.lists = append(this.lists, &elem)
}

func (this NestedInteger) GetList() []*NestedInteger {
	return this.lists
}

type NestedIterator struct {
	index int
	nums  []int
}

func nestedIteratorHelper(nestedList []*NestedInteger) []int {
	var res []int
	for _, node := range nestedList {
		if node.IsInteger() {
			res = append(res, node.GetInteger())
		} else {
			res = append(res, nestedIteratorHelper(node.GetList())...)
		}
	}
	return res
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	return &NestedIterator{
		index: 0,
		nums:  nestedIteratorHelper(nestedList),
	}
}

func (this *NestedIterator) Next() int {
	value := this.nums[this.index]
	this.index += 1
	return value
}

func (this *NestedIterator) HasNext() bool {
	return this.index < len(this.nums)
}
