package question_941_950

import (
	"testing"
)

func Test_minIncrementForUnique(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{[]int{1, 2, 2}, 1},
		{[]int{3, 2, 1, 2, 1, 7}, 6},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minIncrementForUnique(tt.A); got != tt.want {
				t.Errorf("minIncrementForUnique() = %v, want %v", got, tt.want)
			}
		})
	}
}
