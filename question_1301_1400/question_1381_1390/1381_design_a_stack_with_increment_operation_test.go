package question_1361_1370

import (
	"fmt"
	"testing"
)

func TestCustomStack(t *testing.T) {
	this := Constructor(3)
	this.Push(1)
	this.Push(2)
	this.Pop()
	this.Push(2)
	this.Push(3)
	this.Push(4)
	fmt.Printf("%+v\n", this)
	this.Increment(5, 100)
	fmt.Printf("%+v\n", this)
	this.Increment(2, 100)
	fmt.Printf("%+v\n", this)
	this.Pop()
	fmt.Printf("%+v\n", this)
	this.Pop()
	fmt.Printf("%+v\n", this)
	this.Pop()
	fmt.Printf("%+v\n", this)
	this.Pop()
}
