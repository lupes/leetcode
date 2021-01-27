package question_1411_1420

import (
	"testing"
)

func Test_minNumberOfFrogs(t *testing.T) {
	tests := []struct {
		croakOfFrogs string
		want         int
	}{
		{"croakcroak", 1},
		{"crcoakroak", 2},
		{"croakcrook", -1},
		{"croakcroa", -1},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minNumberOfFrogs(tt.croakOfFrogs); got != tt.want {
				t.Errorf("minNumberOfFrogs() = %v, want %v", got, tt.want)
			}
		})
	}
}
