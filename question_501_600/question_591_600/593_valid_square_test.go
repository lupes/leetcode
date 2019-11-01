package question_591_600

import "testing"

func Test_validSquare(t *testing.T) {
	tests := []struct {
		p1   []int
		p2   []int
		p3   []int
		p4   []int
		want bool
	}{
		{[]int{0, 0}, []int{1, 1}, []int{1, 0}, []int{0, 1}, true},
		{[]int{0, 0}, []int{2, 2}, []int{2, 0}, []int{0, 2}, true},
		{[]int{-1, 0}, []int{0, 1}, []int{1, 0}, []int{0, -1}, true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := validSquare(tt.p1, tt.p2, tt.p3, tt.p4); got != tt.want {
				t.Errorf("validSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
