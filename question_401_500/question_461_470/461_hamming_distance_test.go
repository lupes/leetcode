package question_461_470

import "testing"

func Test_hammingDistance(t *testing.T) {
	tests := []struct {
		name string
		x    int
		y    int
		want int
	}{
		{"test", 0, 0, 0},
		{"test", 1, 0, 1},
		{"test", 1, 2, 2},
		{"test", 1, 3, 1},
		{"test", 1, 4, 2},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := hammingDistance(tt.x, tt.y); got != tt.want {
				t.Errorf("hammingDistance() = %v, want %v", got, tt.want)
			}
		})
	}
}
