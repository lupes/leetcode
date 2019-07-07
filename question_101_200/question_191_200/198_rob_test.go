package question_191_200

import "testing"

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test", nil, 0},
		{"test", []int{1}, 1},
		{"test", []int{1, 2}, 2},
		{"test", []int{1, 2, 3}, 4},
		{"test", []int{1, 5, 3}, 5},
		{"test", []int{1, 2, 3, 1}, 4},
		{"test", []int{2, 7, 9, 3, 1}, 12},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
