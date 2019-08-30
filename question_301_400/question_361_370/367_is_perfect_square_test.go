package question_361_370

import "testing"

func Test_isPerfectSquare(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"test", 1, true},
		{"test", 2, false},
		{"test", 16, true},
		{"test", 14, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPerfectSquare(tt.num); got != tt.want {
				t.Errorf("isPerfectSquare() = %v, want %v", got, tt.want)
			}
		})
	}
}
