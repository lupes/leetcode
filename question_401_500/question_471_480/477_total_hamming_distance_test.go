package question_471_480

import (
	"testing"
)

func Test_totalHammingDistance(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{4, 14, 2}, 6},
		{[]int{4, 14, 4}, 4},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := totalHammingDistance(tt.nums); got != tt.want {
				t.Errorf("totalHammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
