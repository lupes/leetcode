package question_1031_1040

import (
	"reflect"
	"testing"
)

func Test_numMovesStones(t *testing.T) {
	tests := []struct {
		a    int
		b    int
		c    int
		want []int
	}{
		{1, 2, 5, []int{1, 2}},
		{4, 3, 2, []int{0, 0}},
		{1, 3, 5, []int{1, 2}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numMovesStones(tt.a, tt.b, tt.c); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numMovesStones() = %v, want %v", got, tt.want)
			}
		})
	}
}
