package question_1051_1060

import (
	"reflect"
	"testing"
)

func Test_rearrangeBarcodes(t *testing.T) {
	tests := []struct {
		barcodes []int
		want     []int
	}{
		{[]int{1, 1, 1, 2, 2, 2, 3}, []int{1, 3, 2, 1, 2, 1, 2}},
		{[]int{1, 1, 1, 2, 2, 2}, []int{1, 2, 1, 2, 1, 2}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := rearrangeBarcodes(tt.barcodes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("rearrangeBarcodes() = %v, want %v", got, tt.want)
			}
		})
	}
}
