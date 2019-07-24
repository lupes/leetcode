package question_261_270

import "testing"

func Test_isUgly(t *testing.T) {
	tests := []struct {
		name string
		num  int
		want bool
	}{
		{"test", 1, true},
		{"test", 2, true},
		{"test", 3, true},
		{"test", 4, true},
		{"test", 5, true},
		{"test", 6, true},
		{"test", 7, false},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := isUgly(tt.num); got != tt.want {
				t.Errorf("isUgly() = %v, want %v", got, tt.want)
			}
		})
	}
}
