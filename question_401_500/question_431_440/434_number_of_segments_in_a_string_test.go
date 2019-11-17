package question_431_440

import "testing"

func Test_countSegments(t *testing.T) {
	tests := []struct {
		name string
		s    string
		want int
	}{
		{"test", "", 0},
		{"test", " ", 0},
		{"test", "  ", 0},
		{"test", "test", 1},
		{"test", "test ", 1},
		{"test", "test test", 2},
		{"test", "Hello, my name is John", 5},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := countSegments(tt.s); got != tt.want {
				t.Errorf("countSegments() = %v, want %v", got, tt.want)
			}
		})
	}
}
