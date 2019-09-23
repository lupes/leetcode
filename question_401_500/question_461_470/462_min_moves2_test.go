package question_461_470

import "testing"

func Test_minMoves2(t *testing.T) {
	tests := []struct {
		name string
		nums []int
		want int
	}{
		{"test", []int{3, 2, 1}, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := minMoves2(tt.nums); got != tt.want {
				t.Errorf("minMoves2() = %v, want %v", got, tt.want)
			}
		})
	}
}
