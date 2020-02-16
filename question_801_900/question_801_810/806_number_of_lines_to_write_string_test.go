package question_801_810

import (
	"reflect"
	"testing"
)

func Test_numberOfLines(t *testing.T) {
	tests := []struct {
		widths []int
		S      string
		want   []int
	}{
		{[]int{10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10},
			"abcdefghijklmnopqrstuvwxyz", []int{3, 60}},
		{[]int{4, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10, 10}, "bbbcccdddaaa", []int{2, 4}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := numberOfLines(tt.widths, tt.S); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("numberOfLines() = %v, want %v", got, tt.want)
			}
		})
	}
}
