package question_621_630

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	queue := Constructor(3)
	if queue.EnQueue(1) != true {
		t.Errorf("%+v\n", queue)
		return
	}
	t.Logf("%+v\n", queue)
	if queue.EnQueue(2) != true {
		t.Errorf("%+v\n", queue)
		return
	}
	t.Logf("%+v\n", queue)
	if queue.EnQueue(3) != true {
		t.Errorf("%+v\n", queue)
		return
	}
	t.Logf("%+v\n", queue)
	if queue.EnQueue(4) != false {
		t.Errorf("%+v\n", queue)
		return
	}
	if queue.Rear() != 3 {
		t.Errorf("%+v\n", queue)
		return
	}
	if queue.IsFull() != true {
		t.Errorf("%+v\n", queue)
		return
	}
	if queue.DeQueue() != true {
		t.Errorf("%+v\n", queue)
		return
	}
	if queue.EnQueue(4) != true {
		t.Errorf("%+v\n", queue)
		return
	}
	if queue.Rear() != 4 {
		t.Errorf("%+v\n", queue)
		return
	}
}
