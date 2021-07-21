package question_861_870

import (
	"reflect"
	"testing"
)

// 23521218 189971378 339517772 341560426 718967141
// 3454610  967482459 978798455 744530040 940238504

func Test_advantageCount(t *testing.T) {
	tests := []struct {
		A    []int
		B    []int
		want []int
	}{
		{[]int{2, 7, 11, 15}, []int{1, 10, 4, 11}, []int{2, 11, 7, 15}},
		{[]int{12, 24, 8, 32}, []int{13, 25, 32, 11}, []int{24, 32, 8, 12}},
		{[]int{718967141, 189971378, 341560426, 23521218, 339517772}, []int{967482459, 978798455, 744530040, 3454610, 940238504}, []int{339517772, 189971378, 718967141, 23521218, 341560426}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := advantageCount(tt.A, tt.B); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("advantageCount() = %v, want %v", got, tt.want)
			}
		})
	}
}
