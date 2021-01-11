package question_1311_1320

import (
	"testing"
)

func Test_minSetSize(t *testing.T) {
	tests := []struct {
		arr  []int
		want int
	}{
		{[]int{3, 3, 3, 3, 5, 5, 5, 2, 2, 7}, 2},
		{[]int{7, 7, 7, 7, 7, 7}, 1},
		{[]int{1, 9}, 1},
		{[]int{1000, 1000, 3, 7}, 1},
		{[]int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minSetSize(tt.arr); got != tt.want {
				t.Errorf("minSetSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
