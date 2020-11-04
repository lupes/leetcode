package question_1331_1340

import (
	"reflect"
	"testing"
)

func Test_sortByBits(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{0, 1, 2, 3, 4, 5, 6, 7, 8}, []int{0, 1, 2, 4, 8, 3, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sortByBits(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sortByBits() = %v, want %v", got, tt.want)
			}
		})
	}
}
