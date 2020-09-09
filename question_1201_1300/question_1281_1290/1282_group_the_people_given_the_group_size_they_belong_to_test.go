package question_1281_1290

import (
	"reflect"
	"testing"
)

func Test_groupThePeople(t *testing.T) {
	tests := []struct {
		groupSizes []int
		want       [][]int
	}{
		{[]int{3, 3, 3, 3, 3, 1, 3}, [][]int{{0, 1, 2}, {5}, {3, 4, 6}}},
		{[]int{2, 1, 3, 3, 3, 2}, [][]int{{1}, {2, 3, 4}, {0, 5}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := groupThePeople(tt.groupSizes); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupThePeople() = %v, want %v", got, tt.want)
			}
		})
	}
}
