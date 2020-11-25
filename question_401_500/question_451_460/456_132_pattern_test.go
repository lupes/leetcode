package question_451_460

import (
	"testing"
)

func Test_find132pattern(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		{[]int{1, 2, 3, 4}, false},
		{[]int{3, 1, 4, 2}, true},
		{[]int{-1, 3, 2, 0}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := find132pattern(tt.nums); got != tt.want {
				t.Errorf("find132pattern() = %v, want %v", got, tt.want)
			}
		})
	}
}
