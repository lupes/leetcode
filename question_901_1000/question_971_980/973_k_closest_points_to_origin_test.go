package question_971_980

import (
	"reflect"
	"testing"
)

func Test_kClosest(t *testing.T) {
	tests := []struct {
		points [][]int
		K      int
		want   [][]int
	}{
		{[][]int{{1, 3}, {-2, 2}}, 1, [][]int{{-2, 2}}},
		{[][]int{{3, 3}, {5, -1}, {-2, 4}}, 2, [][]int{{3, 3}, {-2, 4}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := kClosest(tt.points, tt.K); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("kClosest() = %v, want %v", got, tt.want)
			}
		})
	}
}
