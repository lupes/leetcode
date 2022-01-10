package question_301_310

import (
	"testing"
)

func Test_isAdditiveNumber(t *testing.T) {
	tests := []struct {
		num  string
		want bool
	}{
		{"112358", true},
		{"112", true},
		{"11920", true},
		{"199100199", true},
		{"101", true},
		{"119020", false},
		{"000", true},
		{"1203", false},
		{"199001200", false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isAdditiveNumber(tt.num); got != tt.want {
				t.Errorf("isAdditiveNumber() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_equal(t *testing.T) {
	tests := []struct {
		a    []int
		b    []int
		c    []int
		want bool
	}{
		{[]int{1}, []int{1}, []int{2}, true},
		{[]int{1}, []int{1, 0}, []int{1, 1}, true},
		{[]int{1}, []int{9, 9}, []int{1, 0, 0}, true},
		{[]int{1, 0, 0}, []int{9, 9}, []int{1, 9, 9}, true},
		{[]int{1}, []int{9, 9}, []int{1, 9, 9}, false},
		{[]int{1, 1, 9}, []int{0, 2, 0}, []int{1, 9, 9}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := equal(tt.a, tt.b, tt.c); got != tt.want {
				t.Errorf("equal() = %v, want %v", got, tt.want)
			}
		})
	}
}
