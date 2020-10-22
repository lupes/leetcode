package question_1021_1030

import (
	"reflect"
	"testing"
)

func Test_camelMatch(t *testing.T) {
	tests := []struct {
		queries []string
		pattern string
		want    []bool
	}{
		{[]string{"FooBar", "FooBarTest", "FootBall", "FrameBuffer", "ForceFeedBack"}, "FB", []bool{true, false, true, true, false}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := camelMatch(tt.queries, tt.pattern); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("camelMatch() = %v, want %v", got, tt.want)
			}
		})
	}
}
