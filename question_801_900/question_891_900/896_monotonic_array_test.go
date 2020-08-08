package question_891_900

import (
	"testing"
)

func Test_isMonotonic(t *testing.T) {
	tests := []struct {
		A    []int
		want bool
	}{
		//{[]int{1, 2, 2, 3}, true},
		//{[]int{6, 5, 4, 4}, true},
		//{[]int{1, 3, 2}, false},
		//{[]int{1, 1, 1}, true},
		{[]int{5, 3, 2, 4, 1}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isMonotonic(tt.A); got != tt.want {
				t.Errorf("isMonotonic() = %v, want %v", got, tt.want)
			}
		})
	}
}
