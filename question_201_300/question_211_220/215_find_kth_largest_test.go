package question_211_220

import "testing"

func Test_findKthLargest(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		k    int
		want int
	}{
		{"test", []int{1, 2, 3}, 2, 2},
		{"test", []int{1, 2, 3}, 1, 3},
		{"test", []int{1, 2, 3}, 3, 1},
		{"test", []int{1, 2, 3}, 3, 1},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := findKthLargest(tt.nums, tt.k); got != tt.want {
				t.Errorf("findKthLargest() = %v, want %v", got, tt.want)
			}
		})
	}
}
