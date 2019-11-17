package question_471_480

import "testing"

func Test_makesquare(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want bool
	}{
		{"test", []int{1, 1, 2, 2, 2}, true},
		{"test", []int{10000000000, 10000000000, 10000000000, 10000000000}, true},
		{"test", []int{10000000000, 10000000000, 10000000000}, false},
		{"test", []int{3, 3, 3, 3, 4}, false},
		{"test", []int{2, 2, 2, 2, 2, 6}, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := makesquare(tt.nums); got != tt.want {
				t.Errorf("makesquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
