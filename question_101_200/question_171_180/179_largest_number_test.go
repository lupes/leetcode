package question_171_180

import "testing"

func Test_largestNumber(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want string
	}{
		{"test", []int{0, 0}, "0"},
		{"test", []int{10, 2}, "210"},
		{"test", []int{3, 30, 34, 5, 9}, "9534330"},
		{"test", []int{6, 5, 6, 5, 6, 5, 6, 5, 6, 5}, "6666655555"},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := largestNumber(tt.nums); got != tt.want {
				t.Errorf("largestNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}
