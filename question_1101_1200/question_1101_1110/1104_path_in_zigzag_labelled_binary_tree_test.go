package question_1101_1110

import (
	"reflect"
	"testing"
)

func Test_pathInZigZagTree(t *testing.T) {
	tests := []struct {
		label int
		want  []int
	}{
		{1, []int{1}},
		{14, []int{1, 3, 4, 14}},
		{26, []int{1, 2, 6, 10, 26}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := pathInZigZagTree(tt.label); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("pathInZigZagTree() = %v, want %v", got, tt.want)
			}
		})
	}
}
