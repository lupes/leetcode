package question_81_90

import "testing"

func Test_search(t *testing.T) {
	tests := []struct {
		name   string
		nums   []int
		target int
		want   bool
	}{
		{"test", nil, 0, false},
		{"test", []int{1}, 1, true},
		{"test", []int{2}, 1, false},
		{"test", []int{1, 2}, 1, true},
		{"test", []int{2, 1}, 1, true},
		{"test", []int{3, 4, 5, 1, 2}, 1, true},
		{"test", []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, 2, true},
		{"test", []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, -1, false},
		{"test", []int{3, 4, 5, 6, 7, 8, 9, 0, 1, 2}, 10, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := search(tt.nums, tt.target); got != tt.want {
				t.Errorf("search() = %v, want %v", got, tt.want)
			}
		})
	}
}
