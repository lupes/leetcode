package question_1251_1260

import (
	"reflect"
	"testing"
)

func Test_reconstructMatrix(t *testing.T) {
	tests := []struct {
		upper  int
		lower  int
		colsum []int
		want   [][]int
	}{
		{2, 1, []int{1, 1, 1}, [][]int{{1, 1, 0}, {0, 0, 1}}},
		{2, 3, []int{2, 2, 1, 1}, nil},
		{5, 5, []int{2, 1, 2, 0, 1, 0, 1, 2, 0, 1}, [][]int{{1, 1, 1, 0, 1, 0, 0, 1, 0, 0}, {1, 0, 1, 0, 0, 0, 1, 1, 0, 1}}},
		{9, 2, []int{0, 1, 2, 0, 0, 0, 0, 0, 2, 1, 2, 1, 2}, nil},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reconstructMatrix(tt.upper, tt.lower, tt.colsum); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructMatrix() = %v, want %v", got, tt.want)
			}
		})
	}
}
