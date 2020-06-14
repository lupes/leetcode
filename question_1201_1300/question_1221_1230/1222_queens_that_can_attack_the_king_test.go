package question_1221_1230

import (
	"reflect"
	"sort"
	"testing"
)

func Test_queensAttacktheKing(t *testing.T) {
	tests := []struct {
		queens [][]int
		king   []int
		want   [][]int
	}{
		{[][]int{{0, 1}, {1, 0}, {4, 0}, {0, 4}, {3, 3}, {2, 4}}, []int{0, 0}, [][]int{{0, 1}, {1, 0}, {3, 3}}},
		{[][]int{{0, 0}, {1, 1}, {2, 2}, {3, 4}, {3, 5}, {4, 4}, {4, 5}}, []int{3, 3}, [][]int{{2, 2}, {3, 4}, {4, 4}}},
		{[][]int{{5, 6}, {7, 7}, {2, 1}, {0, 7}, {1, 6}, {5, 1}, {3, 7}, {0, 3}, {4, 0}, {1, 2}, {6, 3}, {5, 0}, {0, 4}, {2, 2}, {1, 1}, {6, 4}, {5, 4}, {0, 0}, {2, 6}, {4, 5}, {5, 2}, {1, 4}, {7, 5}, {2, 3}, {0, 5}, {4, 2}, {1, 0}, {2, 7}, {0, 1}, {4, 6}, {6, 1}, {0, 6}, {4, 3}, {1, 7}},
			[]int{3, 4}, [][]int{{2, 3}, {1, 4}, {1, 6}, {3, 7}, {4, 3}, {5, 4}, {4, 5}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			want := tt.want
			got := queensAttacktheKing(tt.queens, tt.king)
			sort.Slice(got, func(i, j int) bool {
				return (got[i][0] > got[j][0]) || (got[i][0] == got[j][0] && got[i][1] > got[j][1])
			})
			sort.Slice(want, func(i, j int) bool {
				return (want[i][0] > want[j][0]) || (want[i][0] == want[j][0] && want[i][1] > want[j][1])
			})
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("queensAttacktheKing() = %v, want %v", got, tt.want)
			}
		})
	}
}
