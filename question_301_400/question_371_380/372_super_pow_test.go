package question_371_380

import (
	"testing"
)

func Test_superPow(t *testing.T) {
	tests := []struct {
		a    int
		b    []int
		want int
	}{
		{2, []int{3}, 8},
		{2, []int{1, 0}, 1024},
		{2, []int{1, 1}, 711},
		{1, []int{4, 3, 3, 8, 5, 2}, 1},
		{2147483647, []int{2, 0, 0}, 1198},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := superPow(tt.a, tt.b); got != tt.want {
				t.Errorf("superPow() = %v, want %v", got, tt.want)
			}
		})
	}
}
