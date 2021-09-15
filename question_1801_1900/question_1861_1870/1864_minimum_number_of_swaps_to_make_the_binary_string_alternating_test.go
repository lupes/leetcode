package question_1861_1870

import (
	"testing"
)

func Test_minSwapsHelper(t *testing.T) {
	tests := []struct {
		s     string
		order string
		want  int
	}{
		{"111000", "01", 2},
		{"111000", "10", 1},
		{"010", "01", 0},
		{"010", "10", -1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minSwapsHelper(tt.s, tt.order); got != tt.want {
				t.Errorf("minSwapsHelper() = %v, want %v", got, tt.want)
			}
		})
	}
}

func Test_minSwaps(t *testing.T) {
	tests := []struct {
		s    string
		want int
	}{
		{"111000", 1},
		{"010", 0},
		{"1110", -1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minSwaps(tt.s); got != tt.want {
				t.Errorf("minSwaps() = %v, want %v", got, tt.want)
			}
		})
	}
}
