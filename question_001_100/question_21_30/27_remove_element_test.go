package question_0011_0020

import "testing"

func Test_removeElement(t *testing.T) {
	tests := []struct {
		nums []int
		val  int
		want int
	}{
		{nil, 1, 0},
		{[]int{2}, 1, 1},
		{[]int{2}, 2, 0},
		{[]int{1, 2}, 1, 1},
		{[]int{3, 2, 2, 3}, 3, 2},
		{[]int{0, 1, 2, 2, 3, 0, 4, 2}, 2, 5},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := removeElement(tt.nums, tt.val); got != tt.want {
				t.Errorf("removeElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
