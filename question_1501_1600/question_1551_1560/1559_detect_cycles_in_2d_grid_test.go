package question_1551_1560

import (
	"testing"
)

func Test_containsCycle(t *testing.T) {
	tests := []struct {
		grid [][]byte
		want bool
	}{
		{[][]byte{{'a', 'a', 'a', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'b', 'b', 'a'}, {'a', 'a', 'a', 'a'}}, true},
		{[][]byte{{'c', 'c', 'c', 'a'}, {'c', 'd', 'c', 'c'}, {'c', 'c', 'e', 'c'}, {'f', 'c', 'c', 'c'}}, true},
		{[][]byte{{'a', 'b', 'b'}, {'b', 'z', 'b'}, {'b', 'b', 'a'}}, false},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := containsCycle(tt.grid); got != tt.want {
				t.Errorf("containsCycle() = %v, want %v", got, tt.want)
			}
		})
	}
}
