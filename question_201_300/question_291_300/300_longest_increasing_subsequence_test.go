package question_291_300

import (
	"testing"
)

func Test_lengthOfLIS(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
		{[]int{10, 9, 2, 5, 3, 7, 101, 18}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := lengthOfLIS(tt.nums); got != tt.want {
				t.Errorf("lengthOfLIS() = %v, want %v", got, tt.want)
			}
		})
	}
}
