package question_1291_1300

import (
	"testing"
)

func Test_findBestValue(t *testing.T) {
	tests := []struct {
		arr    []int
		target int
		want   int
	}{
		{[]int{4, 9, 3}, 10, 3},
		{[]int{2, 3, 5}, 10, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findBestValue(tt.arr, tt.target); got != tt.want {
				t.Errorf("findBestValue() = %v, want %v", got, tt.want)
			}
		})
	}
}
