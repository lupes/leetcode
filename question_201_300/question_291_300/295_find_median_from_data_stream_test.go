package question_291_300

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	this := Constructor()

	this.AddNum(1)
	if got, want := this.FindMedian(), 1.0; got != want {
		t.Errorf("FindMedian() = %v, want %v", got, want)
	}

	this.AddNum(2)
	if got, want := this.FindMedian(), 1.5; got != want {
		t.Errorf("FindMedian() = %v, want %v", got, want)
	}

	this.AddNum(3)
	if got, want := this.FindMedian(), 2.0; got != want {
		t.Errorf("FindMedian() = %v, want %v", got, want)
	}
}
