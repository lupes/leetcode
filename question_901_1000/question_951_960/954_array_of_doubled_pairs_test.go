package question_951_960

import (
	"testing"
)

func Test_canReorderDoubled(t *testing.T) {
	tests := []struct {
		A    []int
		want bool
	}{
		{[]int{3, 1, 3, 6}, false},
		{[]int{2, 1, 2, 6}, false},
		{[]int{2, 1, 3, 6}, true},
		{[]int{4, -2, 2, -4}, true},
		{[]int{1, 2, 4, 16, 8, 4}, false},
		{[]int{1, 2, 4, 16, 8, 8}, true},
		{[]int{1, 2, 2, 4, 16, 8}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canReorderDoubled(tt.A); got != tt.want {
				t.Errorf("canReorderDoubled() = %v, want %v", got, tt.want)
			}
		})
	}
}
