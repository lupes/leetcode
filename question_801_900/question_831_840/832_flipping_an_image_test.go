package question_831_840

import (
	"reflect"
	"testing"
)

func Test_flipAndInvertImage(t *testing.T) {
	tests := []struct {
		A    [][]int
		want [][]int
	}{
		{[][]int{{1, 1, 0}, {1, 0, 1}, {0, 0, 0}}, [][]int{{1, 0, 0}, {0, 1, 0}, {1, 1, 1}}},
		{[][]int{{1, 1, 0, 0}, {1, 0, 0, 1}, {0, 1, 1, 1}, {1, 0, 1, 0}},
			[][]int{{1, 1, 0, 0}, {0, 1, 1, 0}, {0, 0, 0, 1}, {1, 0, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := flipAndInvertImage(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("flipAndInvertImage() = %v, want %v", got, tt.want)
			}
		})
	}
}
