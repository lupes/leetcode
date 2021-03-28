package question_1551_1560

import (
	"reflect"
	"testing"
)

func Test_mostVisited(t *testing.T) {
	tests := []struct {
		n      int
		rounds []int
		want   []int
	}{
		{4, []int{1, 3, 1, 2}, []int{1, 2}},
		{2, []int{2, 1, 2, 1, 2, 1, 2, 1, 2}, []int{2}},
		{7, []int{1, 3, 5, 7}, []int{1, 2, 3, 4, 5, 6, 7}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := mostVisited(tt.n, tt.rounds)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("mostVisited() = %v, want %v", got, tt.want)
			}
		})
	}
}
