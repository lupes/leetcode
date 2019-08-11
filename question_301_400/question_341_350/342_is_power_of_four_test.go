package question_341_350

import "testing"

func Test_isPowerOfFour(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"test", 1, true},
		{"test", 16, true},
		{"test", 5, false},
		{"test", 32, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isPowerOfFour(tt.num); got != tt.want {
				t.Errorf("isPowerOfFour() = %v, want %v", got, tt.want)
			}
		})
	}
}
