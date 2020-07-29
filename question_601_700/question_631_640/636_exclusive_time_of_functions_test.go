package question_631_640

import (
	"reflect"
	"testing"
)

func Test_exclusiveTime(t *testing.T) {
	tests := []struct {
		n    int
		logs []string
		want []int
	}{
		{2, []string{"0:start:0", "1:start:2", "1:end:5", "0:end:6"}, []int{3, 4}},
		{1, []string{"0:start:0", "0:start:1", "0:start:2", "0:end:3", "0:end:4", "0:end:5"}, []int{6}},
		{3, []string{"0:start:0", "0:start:1", "0:end:2", "1:start:3", "2:start:4", "2:end:5", "1:end:6", "0:end:7"}, []int{4, 2, 2}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := exclusiveTime(tt.n, tt.logs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("exclusiveTime() = %v, want %v", got, tt.want)
			}
		})
	}
}
