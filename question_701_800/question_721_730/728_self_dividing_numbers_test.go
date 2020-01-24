package question_721_730

import (
	"reflect"
	"testing"
)

func Test_selfDividingNumbers(t *testing.T) {
	tests := []struct {
		left  int
		right int
		want  []int
	}{
		{1, 22, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
		{1, 10000, []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 11, 12, 15, 22}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := selfDividingNumbers(tt.left, tt.right); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("selfDividingNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
