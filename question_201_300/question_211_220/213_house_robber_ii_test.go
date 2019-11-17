package question_211_220

import "testing"

func Test_rob(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test", []int{2}, 2},
		{"test", []int{2, 3}, 3},
		{"test", []int{2, 3, 2}, 3},
		{"test", []int{1, 2, 3, 1}, 4},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := rob(tt.nums); got != tt.want {
				t.Errorf("rob() = %v, want %v", got, tt.want)
			}
		})
	}
}
