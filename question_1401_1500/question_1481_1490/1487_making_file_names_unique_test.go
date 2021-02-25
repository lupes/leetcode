package question_1481_1490

import (
	"reflect"
	"testing"
)

func Test_getFolderNames(t *testing.T) {
	tests := []struct {
		names []string
		want  []string
	}{
		{[]string{"pes", "fifa", "gta", "pes(2019)"}, []string{"pes", "fifa", "gta", "pes(2019)"}},
		{[]string{"gta", "gta(1)", "gta", "avalon"}, []string{"gta", "gta(1)", "gta(2)", "avalon"}},
		{[]string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece"}, []string{"onepiece", "onepiece(1)", "onepiece(2)", "onepiece(3)", "onepiece(4)"}},
		{[]string{"wano", "wano", "wano", "wano"}, []string{"wano", "wano(1)", "wano(2)", "wano(3)"}},
		{[]string{"kaido", "kaido(1)", "kaido", "kaido(1)", "kaido(2)"}, []string{}},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := getFolderNames(tt.names); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getFolderNames() = %#v, want %#v", got, tt.want)
			}
		})
	}
}
