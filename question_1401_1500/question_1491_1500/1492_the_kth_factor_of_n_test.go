package question_1491_1500

import (
	"testing"
)

func Test_kthFactor(t *testing.T) {
	tests := []struct {
		n    int
		k    int
		want int
	}{
		{12, 3, 3},
		{7, 2, 7},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := kthFactor(tt.n, tt.k); got != tt.want {
				t.Errorf("kthFactor() = %v, want %v", got, tt.want)
			}
		})
	}
}
