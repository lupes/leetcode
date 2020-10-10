package question_281_290

import (
	"testing"
)

func TestPeekingIterator(t *testing.T) {
	iter := &Iterator{arr: []int{1, 2, 3}}
	peek := Constructor(iter)
	if got, want := peek.next(), 1; got != want {
		t.Errorf("next got=%v want=%v\n", got, want)
	}

	if got, want := peek.peek(), 2; got != want {
		t.Errorf("peek got=%v want=%v\n", got, want)
	}

	if got, want := peek.next(), 2; got != want {
		t.Errorf("next got=%v want=%v\n", got, want)
	}

	if got, want := peek.next(), 3; got != want {
		t.Errorf("next got=%v want=%v\n", got, want)
	}

	if got, want := peek.hasNext(), false; got != want {
		t.Errorf("next got=%v want=%v\n", got, want)
	}
}
