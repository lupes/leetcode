package question_1631_1640

import (
	"testing"
)

func Test_canFormArray(t *testing.T) {
	tests := []struct {
		arr    []int
		pieces [][]int
		want   bool
	}{
		{[]int{85}, [][]int{{85}}, true},
		{[]int{15, 88}, [][]int{{88}, {15}}, true},
		{[]int{49, 18, 16}, [][]int{{16, 18, 49}}, false},
		{[]int{91, 4, 64, 78}, [][]int{{78}, {4, 64}, {91}}, true},
		{[]int{1, 3, 5, 7}, [][]int{{2, 4, 6, 8}}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canFormArray(tt.arr, tt.pieces); got != tt.want {
				t.Errorf("canFormArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
