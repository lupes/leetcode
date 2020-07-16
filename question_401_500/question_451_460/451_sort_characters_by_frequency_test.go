package question_451_460

import (
	"testing"
)

func Test_frequencySort(t *testing.T) {
	tests := []struct {
		s    string
		want string
	}{
		{"tree", "eert"},
		{"cccaaa", "aaaccc"},
		{"Aabb", "bbaA"},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := frequencySort(tt.s); got != tt.want {
				t.Errorf("frequencySort() = %v, want %v", got, tt.want)
			}
		})
	}
}
