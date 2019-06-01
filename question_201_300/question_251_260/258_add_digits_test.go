package question_251_260

import "testing"

func Test_addDigits(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want int
	}{
		{"test", 1, 1},
		{"test", 10, 1},
		{"test", 38, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := addDigits(tt.num); got != tt.want {
				t.Errorf("addDigits() = %v, want %v", got, tt.want)
			}
		})
	}
}
