package question_821_830

import (
	"reflect"
	"testing"
)

func Test_shortestToChar(t *testing.T) {
	tests := []struct {
		S    string
		C    byte
		want []int
	}{
		{"loveleetcode", 'e', []int{3, 2, 1, 0, 1, 0, 0, 1, 2, 2, 1, 0}},
		{"l", 'l', []int{0}},
		{"lo", 'l', []int{0, 1}},
		{"lo", 'o', []int{1, 0}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := shortestToChar(tt.S, tt.C); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("shortestToChar() = %v, want %v", got, tt.want)
			}
		})
	}
}
