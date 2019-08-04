package question_281_290

import (
	"reflect"
	"testing"
)

func Test_gameOfLife(t *testing.T) {
	tests := []struct {
		name  string
		board [][]int
		want  [][]int
	}{
		{"test", [][]int{{0, 1, 0}, {0, 0, 1}, {1, 1, 1}, {0, 0, 0}}, [][]int{{0, 0, 0}, {1, 0, 1}, {0, 1, 1}, {0, 1, 0}}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			gameOfLife(tt.board)
			if !reflect.DeepEqual(tt.board, tt.want) {
				t.Errorf("gameOfLife got=%v want=%v\n", tt.board, tt.want)
			}
		})
	}
}
