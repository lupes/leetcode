package question_531_540

import "testing"

func Test_singleNonDuplicate(t *testing.T) {
	tests := []struct {
		nums []int
		want int
	}{
		{[]int{3}, 3},
		{[]int{3, 3, 7}, 7},
		{[]int{1, 1, 2, 3, 3, 4, 4, 8, 8}, 2},
		{[]int{3, 3, 7, 7, 10, 11, 11}, 10},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := singleNonDuplicate(tt.nums); got != tt.want {
				t.Errorf("singleNonDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
