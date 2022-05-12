package question_941_950

import (
	"testing"
)

func Test_minDeletionSize(t *testing.T) {

	tests := []struct {
		A    []string
		want int
	}{
		{[]string{"abc", "bce", "cae"}, 1},
		{[]string{"cba", "daf", "ghi"}, 1},
		{[]string{"a", "b"}, 0},
		{[]string{"zyx", "wvu", "tsr"}, 3},
		{[]string{"rrjk", "furt", "guzm"}, 2},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := minDeletionSize(tt.A); got != tt.want {
				t.Errorf("minDeletionSize() = %v, want %v", got, tt.want)
			}
		})
	}
}
