package question_161_170

import "testing"

func Test_majorityElement(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test", []int{3}, 3},
		{"test", []int{3, 2, 3}, 3},
		{"test", []int{3, 2, 3, 3}, 3},
		{"test", []int{3, 2, 2, 3, 3}, 3},
		{"test", []int{2, 2, 1, 1, 1, 2, 2}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := majorityElement(tt.nums); got != tt.want {
				t.Errorf("majorityElement() = %v, want %v", got, tt.want)
			}
		})
	}
}
