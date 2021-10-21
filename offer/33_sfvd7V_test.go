package offer

import (
	"reflect"
	"testing"
)

func Test_groupAnagrams(t *testing.T) {
	tests := []struct {
		strs []string
		want [][]string
	}{
		{[]string{"eat", "tea", "tan", "ate", "nat", "bat"}, [][]string{{"eat", "tea", "ate"}, {"tan", "nat"}, {"bat"}}},
		{[]string{""}, [][]string{{""}}},
		{[]string{"a"}, [][]string{{"a"}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := groupAnagrams(tt.strs); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("groupAnagrams() = %#v, want %v", got, tt.want)
			}
		})
	}
}
