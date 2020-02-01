package question_731_740

import (
	"reflect"
	"testing"
)

func Test_floodFill(t *testing.T) {
	tests := []struct {
		image    [][]int
		sr       int
		sc       int
		newColor int
		want     [][]int
	}{
		{[][]int{{1, 1, 1}, {1, 1, 0}, {1, 0, 1}}, 1, 1, 2,
			[][]int{{2, 2, 2}, {2, 2, 0}, {2, 0, 1}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := floodFill(tt.image, tt.sr, tt.sc, tt.newColor); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("floodFill() = %v, want %v", got, tt.want)
			}
		})
	}
}
