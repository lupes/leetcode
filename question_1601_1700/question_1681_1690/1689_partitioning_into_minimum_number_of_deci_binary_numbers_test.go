package question_1681_1690

import (
	"testing"
)

func Test_minPartitions(t *testing.T) {
	tests := []struct {
		n    string
		want int
	}{
		{"32", 3},
		{"82734", 8},
		{"27346209830709182346", 9},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minPartitions(tt.n); got != tt.want {
				t.Errorf("minPartitions() = %v, want %v", got, tt.want)
			}
		})
	}
}
