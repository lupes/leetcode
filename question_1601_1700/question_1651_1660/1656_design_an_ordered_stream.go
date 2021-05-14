package question_1651_1660

// 1656. 设计有序流
// https://leetcode-cn.com/problems/design-an-ordered-stream/
// Topics: 设计 数组

type OrderedStream struct {
	Data []string
	Id   int
}

func Constructor(n int) OrderedStream {
	return OrderedStream{
		Data: make([]string, n+2),
		Id:   1,
	}
}

func (this *OrderedStream) Insert(idKey int, value string) []string {
	this.Data[idKey] = value
	var start = this.Id
	for len(this.Data[this.Id]) > 0 {
		this.Id++
	}
	return this.Data[start:this.Id]
}
