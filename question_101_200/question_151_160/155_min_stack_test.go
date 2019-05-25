package question_151_160

import (
	"reflect"
	"testing"
)

func TestStack(t *testing.T) {
	t.Run("test", func(t *testing.T) {
		var minStack = Constructor()
		minStack.Push(-2)
		minStack.Push(0)
		minStack.Push(-3)
		if minStack.GetMin() != -3 {
			t.Errorf("error")
		}
		minStack.Push(1)
		minStack.Push(-1)
		minStack.Pop()
		if reflect.DeepEqual(minStack.sorts, []int{}) || reflect.DeepEqual(minStack.sorts, []int{}) {
			t.Errorf("error")
			return
		}
		minStack.Pop()
		if minStack.Top() == 0 {
			t.Errorf("error")
			return
		}
		if minStack.GetMin() == -2 {
			t.Errorf("error")
			return
		}
		minStack.Pop()
		minStack.Pop()
		minStack.Pop()
		minStack.Pop()
	})
	t.Run("test", func(t *testing.T) {
		var minStack = Constructor()
		minStack.Push(512)
		minStack.Push(-1024)
		minStack.Push(-1024)
		minStack.Push(512)
		minStack.Pop()
		minStack.GetMin()
		minStack.Pop()
		minStack.GetMin()
		minStack.Pop()
		minStack.GetMin()
	})
}
