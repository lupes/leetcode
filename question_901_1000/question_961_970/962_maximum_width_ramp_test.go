package question_961_970

import (
	"testing"
)

func Test_maxWidthRamp(t *testing.T) {
	tests := []struct {
		A    []int
		want int
	}{
		//{[]int{7, 2, 5, 4}, 2},
		//{[]int{6, 0, 8, 2, 1, 5}, 4},
		//{[]int{9, 8, 1, 0, 1, 9, 4, 0, 4, 1}, 7},
		{[]int{10, 10, 10, 7, 1, 6, 2, 1, 7}, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxWidthRamp(tt.A); got != tt.want {
				t.Errorf("maxWidthRamp() = %v, want %v", got, tt.want)
			}
		})
	}
}
