package question_1451_1460

import (
	"testing"
)

func Test_canBeEqual(t *testing.T) {
	tests := []struct {
		target []int
		arr    []int
		want   bool
	}{
		{[]int{1, 2, 3, 4}, []int{2, 4, 1, 3}, true},
		{[]int{7}, []int{7}, true},
		{[]int{1, 12}, []int{12, 1}, true},
		{[]int{3, 7, 9}, []int{3, 7, 11}, false},
		{[]int{1, 1, 1, 1, 1}, []int{1, 1, 1, 1, 1}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := canBeEqual(tt.target, tt.arr); got != tt.want {
				t.Errorf("canBeEqual() = %v, want %v", got, tt.want)
			}
		})
	}
}
