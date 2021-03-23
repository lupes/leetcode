package question_1551_1560

import (
	"testing"
)

func Test_minOperations(t *testing.T) {
	tests := []struct {
		n    int
		want int
	}{
		{3, 2},
		{6, 9},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minOperations(tt.n); got != tt.want {
				t.Errorf("minOperations() = %v, want %v", got, tt.want)
			}
		})
	}
}
