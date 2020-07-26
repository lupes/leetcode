package question_621_630

import (
	"testing"
)

func Test_leastInterval(t *testing.T) {
	tests := []struct {
		tasks []byte
		n     int
		want  int
	}{
		{[]byte{'A', 'A', 'A', 'B', 'B', 'B'}, 2, 8},
		{[]byte{'A', 'A', 'A', 'B', 'B'}, 2, 7},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'C'}, 2, 7},
		{[]byte{'A', 'A', 'A', 'B', 'B', 'C', 'D'}, 2, 7},
		{[]byte{'A', 'A', 'B', 'B', 'C', 'C', 'D', 'D', 'E', 'E', 'F', 'F', 'G', 'G', 'H', 'H', 'I', 'I', 'J', 'J', 'K', 'K', 'L', 'L', 'M', 'M', 'N', 'N', 'O', 'O', 'P', 'P', 'Q', 'Q', 'R', 'R', 'S', 'S', 'T', 'T', 'U', 'U', 'V', 'V', 'W', 'W', 'X', 'X', 'Y', 'Y', 'Z', 'Z'}, 2, 52},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := leastInterval(tt.tasks, tt.n); got != tt.want {
				t.Errorf("leastInterval() = %v, want %v", got, tt.want)
			}
		})
	}
}
