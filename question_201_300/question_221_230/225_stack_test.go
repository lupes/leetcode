package question_221_230

import (
	"testing"
)

func TestStack(t *testing.T) {
	obj := Constructor()
	obj.Push(1)
	param3 := obj.Top()
	if param3 != 1 {
		t.Errorf("got=%d want=%d\n", param3, 1)
	}

	param2 := obj.Empty()
	if param2 != false {
		t.Errorf("got=%v want=%v\n", param2, false)
	}
	param1 := obj.Pop()
	if param1 != 1 {
		t.Errorf("got=%d want=%d\n", param1, 1)
	}

	param4 := obj.Empty()
	if param4 != true {
		t.Errorf("got=%v want=%v\n", param4, true)
	}
}
