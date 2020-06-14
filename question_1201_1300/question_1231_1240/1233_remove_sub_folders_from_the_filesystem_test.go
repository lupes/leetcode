package question_1231_1240

import (
	"reflect"
	"sort"
	"testing"
)

func Test_removeSubfolders(t *testing.T) {
	tests := []struct {
		folder []string
		want   []string
	}{
		{[]string{"/a", "/a/b", "/c/d", "/c/d/e", "/c/f"}, []string{"/a", "/c/d", "/c/f"}},
		{[]string{"/a", "/a/b/c", "/a/b/d"}, []string{"/a"}},
		{[]string{"/a/b/c", "/a/b/d", "/a/b/ca"}, []string{"/a/b/c", "/a/b/ca", "/a/b/d"}},
		{[]string{"/ah/al/am", "/ah/al"}, []string{"/ah/al"}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			got := removeSubfolders(tt.folder)
			sort.Strings(got)
			sort.Strings(tt.want)
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("removeSubfolders() = %v, want %v", got, tt.want)
			}
		})
	}
}
