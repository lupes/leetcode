package question_941_950

import (
	"testing"
)

func Test_validateStackSequences(t *testing.T) {
	tests := []struct {
		pushed []int
		popped []int
		want   bool
	}{
		{[]int{1, 2, 3, 4, 5}, []int{4, 5, 3, 2, 1}, true},
		{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 2, 1}, true},
		{[]int{1, 2, 3, 4, 5}, []int{4, 3, 5, 1, 2}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := validateStackSequences(tt.pushed, tt.popped); got != tt.want {
				t.Errorf("validateStackSequences() = %v, want %v", got, tt.want)
			}
		})
	}
}
