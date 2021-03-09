package question_1521_1530

import (
	"testing"
)

func Test_numOfSubarrays(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{1, 3, 5}, 4},
		{[]int{2, 4, 6}, 0},
		{[]int{1, 2, 3, 4, 5, 6, 7}, 16},
		{[]int{100, 100, 99, 99}, 4},
		{[]int{7}, 1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numOfSubarrays(tt.arr); got != tt.want {
				t.Errorf("numOfSubarrays() = %v, want %v", got, tt.want)
			}
		})
	}
}
