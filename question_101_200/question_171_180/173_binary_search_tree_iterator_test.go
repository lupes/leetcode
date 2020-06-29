package question_171_180

import (
	"math"
	"testing"

	"github.com/lupes/leetcode/common"
)

func TestBSTIterator_Next(t *testing.T) {
	itarator := Constructor(common.NewNodeV2(7, 3, 15, math.MaxInt32, math.MaxInt32, 9, 20))
	if got, want := itarator.Next(), 3; got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	if got, want := itarator.Next(), 7; got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	if got, want := itarator.HasNext(), true; got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	if got, want := itarator.Next(), 9; got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	if got, want := itarator.HasNext(), true; got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	if got, want := itarator.Next(), 15; got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	if got, want := itarator.HasNext(), true; got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	if got, want := itarator.Next(), 20; got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
	if got, want := itarator.HasNext(), false; got != want {
		t.Errorf("Next() = %v, want %v", got, want)
	}
}
