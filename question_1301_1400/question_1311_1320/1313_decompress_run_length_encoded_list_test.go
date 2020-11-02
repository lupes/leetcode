package question_1311_1320

import (
	"reflect"
	"testing"
)

func Test_decompressRLElist(t *testing.T) {
	tests := []struct {
		nums []int
		want []int
	}{
		{[]int{1, 2, 3, 4}, []int{2, 4, 4, 4}},
		{[]int{1, 1, 2, 3}, []int{1, 3, 3}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := decompressRLElist(tt.nums); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("decompressRLElist() = %v, want %v", got, tt.want)
			}
		})
	}
}
