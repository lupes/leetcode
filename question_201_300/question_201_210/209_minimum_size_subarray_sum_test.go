package question_201_210

import "testing"

func Test_minSubArrayLen(t *testing.T) {
	tests := []struct {
		name string
		s    int
		nums []int
		want int
	}{
		{"test", 1, nil, 0},
		{"test", 1, []int{1}, 1},
		{"test", 2, []int{1}, 0},
		{"test", 2, []int{1, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minSubArrayLen(tt.s, tt.nums); got != tt.want {
				t.Errorf("minSubArrayLen() = %v, want %v", got, tt.want)
			}
		})
	}
}
