package question_521_530

import (
	"testing"
)

func Test_checkSubarraySum(t *testing.T) {
	tests := []struct {
		nums []int
		k    int
		want bool
	}{
		{[]int{23, 2}, 0, false},
		{[]int{0, 0}, 0, true},
		{[]int{23, 2, 4, 6, 7}, 6, true},
		{[]int{23, 2, 4}, 6, true},
		{[]int{23, 2}, 6, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := checkSubarraySum(tt.nums, tt.k); got != tt.want {
				t.Errorf("checkSubarraySum() = %v, want %v", got, tt.want)
			}
		})
	}
}
