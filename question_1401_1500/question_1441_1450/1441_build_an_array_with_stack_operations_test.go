package question_1441_1450

import (
	"reflect"
	"testing"
)

func Test_buildArray(t *testing.T) {
	tests := []struct {
		target []int
		n      int
		want   []string
	}{
		{[]int{1, 3}, 3, []string{"Push", "Push", "Pop", "Push"}},
		{[]int{1, 2, 3}, 3, []string{"Push", "Push", "Push"}},
		{[]int{1, 2}, 4, []string{"Push", "Push"}},
		{[]int{2, 3, 4}, 4, []string{"Push", "Pop", "Push", "Push", "Push"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := buildArray(tt.target, tt.n); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("buildArray() = %v, want %v", got, tt.want)
			}
		})
	}
}
