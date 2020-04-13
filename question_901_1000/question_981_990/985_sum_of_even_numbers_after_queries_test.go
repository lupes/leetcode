package question_981_990

import (
	"reflect"
	"testing"
)

func Test_sumEvenAfterQueries(t *testing.T) {
	tests := []struct {
		A       []int
		queries [][]int
		want    []int
	}{
		{[]int{1, 2, 3, 4}, [][]int{{1, 0}, {-3, 1}, {-4, 0}, {2, 3}}, []int{8, 6, 2, 4}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := sumEvenAfterQueries(tt.A, tt.queries); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("sumEvenAfterQueries() = %v, want %v", got, tt.want)
			}
		})
	}
}
