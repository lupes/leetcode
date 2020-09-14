package question_1281_1290

import (
	"fmt"
)

// 1286. 字母组合迭代器
// https://leetcode-cn.com/problems/iterator-for-combination/
// Topics: 回溯算法

type CombinationIterator struct {
	chars  []byte
	flag   []int
	length int
	t      bool
}

func (this CombinationIterator) String() string {
	return fmt.Sprintf("len:%d flag:%+v\n", this.length, this.flag)
}

func Constructor(characters string, combinationLength int) CombinationIterator {
	this := CombinationIterator{
		chars:  []byte(characters),
		flag:   make([]int, combinationLength),
		length: combinationLength,
		t:      true,
	}
	for i := 0; i < this.length; i++ {
		this.flag[i] = i
	}
	return this
}

func (this *CombinationIterator) Next() string {
	var res = make([]byte, this.length)
	for i, f := range this.flag {
		res[i] = this.chars[f]
	}
	if this.flag[0] == len(this.chars)-this.length {
		this.t = false
		return string(res)
	}
	for i := this.length - 1; i >= 0; i-- {
		if this.flag[i] < len(this.chars)-(this.length-i) {
			this.flag[i]++
			for t := i + 1; t < this.length; t++ {
				this.flag[t] = this.flag[i] + (t - i)
			}
			break
		}
	}
	return string(res)
}

func (this *CombinationIterator) HasNext() bool {
	return this.t
}
