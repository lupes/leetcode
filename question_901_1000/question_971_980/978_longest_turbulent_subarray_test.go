package question_971_980

import (
	"testing"
)

func Test_maxTurbulenceSize(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		{[]int{9, 4, 2, 10, 7, 8, 8, 1, 9}, 5},
		{[]int{4, 8, 12, 16}, 2},
		{[]int{4}, 1},
		{[]int{1, 4}, 2},
		{[]int{4, 4}, 1},
		{[]int{4, 4, 4}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxTurbulenceSize(tt.A); got != tt.want {
				t.Errorf("maxTurbulenceSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
