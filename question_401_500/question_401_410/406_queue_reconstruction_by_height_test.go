package question_401_410

import (
	"reflect"
	"testing"
)

func Test_reconstructQueue(t *testing.T) {
	tests := []struct {
		people [][]int
		want   [][]int
	}{
		{[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}},
			[][]int{{7, 0}, {4, 4}, {7, 1}, {5, 0}, {6, 1}, {5, 2}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := reconstructQueue(tt.people); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("reconstructQueue() = %v, want %v", got, tt.want)
			}
		})
	}
}
