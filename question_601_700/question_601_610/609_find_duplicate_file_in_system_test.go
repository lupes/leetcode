package question_601_610

import (
	"reflect"
	"testing"
)

func Test_findDuplicate(t *testing.T) {
	tests := []struct {
		paths []string
		want  [][]string
	}{
		{[]string{"root/a 1.txt(abcd) 2.txt(efgh)", "root/c 3.txt(abcd)", "root/c/d 4.txt(efgh)", "root 4.txt(efgh)"},
			[][]string{{"root/a/1.txt", "root/c/3.txt"}, {"root/a/2.txt", "root/c/d/4.txt", "root/4.txt"}}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := findDuplicate(tt.paths); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("findDuplicate() = %v, want %v", got, tt.want)
			}
		})
	}
}
