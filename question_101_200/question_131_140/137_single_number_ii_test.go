package question_131_140

import "testing"

func Test_singleNumber2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test", []int{2}, 2},
		{"test", []int{2, 2, 3, 2}, 3},
		{"test", []int{0, 1, 0, 1, 0, 1, 99}, 99},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := singleNumber2(tt.nums); got != tt.want {
				t.Errorf("singleNumber2() = %v, want %v", got, tt.want)
			}
		})
	}
}
