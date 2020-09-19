package question_1291_1300

import (
	"reflect"
	"testing"
)

func Test_replaceElements(t *testing.T) {
	tests := []struct {
		arr  []int
		want []int
	}{
		{[]int{17, 18, 5, 4, 6, 1}, []int{18, 6, 6, 6, 1, -1}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := replaceElements(tt.arr); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("replaceElements() = %v, want %v", got, tt.want)
			}
		})
	}
}
