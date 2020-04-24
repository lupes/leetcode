package question_1001_1010

import (
	"reflect"
	"testing"
)

func Test_commonChars(t *testing.T) {
	tests := []struct {
		A    []string
		want []string
	}{
		{[]string{"bella", "label", "roller"}, []string{"e", "l", "l"}},
		{[]string{"cool", "lock", "cook"}, []string{"c", "o"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := commonChars(tt.A); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("commonChars() = %v, want %v", got, tt.want)
			}
		})
	}
}
