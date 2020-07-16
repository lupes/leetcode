package question_451_460

import (
	"bytes"
	"fmt"
)

// 451. 根据字符出现频率排序
// https://leetcode-cn.com/problems/sort-characters-by-frequency
// Topics: 堆 哈希表

type node struct {
	c    int32
	f    int
	next *node
}

func (n *node) String() string {
	return fmt.Sprintf("%s%d(%s)", string(n.c), n.f, n.next)
}

func frequencySort(s string) string {
	var flag = make(map[int32]int)
	for _, c := range s {
		flag[c]++
	}
	var head = &node{}
	for k, v := range flag {
		var tmp = head
		for ; tmp.next != nil && v < tmp.next.f; tmp = tmp.next {
		}
		n := &node{
			c:    k,
			f:    v,
			next: tmp.next,
		}
		tmp.next = n
	}

	buffer := bytes.NewBuffer(make([]byte, 0, len(s)))
	for n := head.next; n != nil; n = n.next {
		for i := 0; i < n.f; i++ {
			buffer.WriteByte(byte(n.c))
		}
	}

	return buffer.String()
}
