package question_1461_1470

import (
	"testing"
)

func Test_maxArea(t *testing.T) {
	tests := []struct {
		h              int
		w              int
		horizontalCuts []int
		verticalCuts   []int
		want           int
	}{
		{5, 4, []int{1, 2, 4}, []int{1, 3}, 4},
		{5, 4, []int{3, 1}, []int{1}, 6},
		{5, 4, []int{3}, []int{3}, 9},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := maxArea(tt.h, tt.w, tt.horizontalCuts, tt.verticalCuts); got != tt.want {
				t.Errorf("maxArea() = %v, want %v", got, tt.want)
			}
		})
	}
}
