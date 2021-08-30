package question_451_460

import (
	"testing"
)

func Test_circularArrayLoop(t *testing.T) {
	tests := []struct {
		nums []int
		want bool
	}{
		//{[]int{2, -1, 1, 2, 2}, true},
		{[]int{-1, 2}, false},
		{[]int{-2, 1, -1, -2, -2}, false},
		{[]int{2, 2, 2, 2, 2, 4, 7}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := circularArrayLoop(tt.nums); got != tt.want {
				t.Errorf("circularArrayLoop() = %v, want %v", got, tt.want)
			}
		})
	}
}
