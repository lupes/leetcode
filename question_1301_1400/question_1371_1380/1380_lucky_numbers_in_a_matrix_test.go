package question_1361_1370

import (
	"reflect"
	"testing"
)

func Test_luckyNumbers(t *testing.T) {
	tests := []struct {
		matrix [][]int
		want   []int
	}{
		{[][]int{{3, 7, 8}, {9, 11, 13}, {15, 16, 17}}, []int{15}},
		{[][]int{{1, 10, 4, 2}, {9, 3, 8, 7}, {15, 16, 17, 12}}, []int{12}},
		{[][]int{{7, 8}, {1, 2}}, []int{7}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := luckyNumbers(tt.matrix); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("luckyNumbers() = %v, want %v", got, tt.want)
			}
		})
	}
}
