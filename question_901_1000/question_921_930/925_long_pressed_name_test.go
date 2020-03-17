package question_921_930

import (
	"testing"
)

func Test_isLongPressedName(t *testing.T) {
	tests := []struct {
		name  string
		typed string
		want  bool
	}{
		{"alex", "aaleex", true},
		{"saeed", "ssaaedd", false},
		{"leelee", "lleeelee", true},
		{"laiden", "laiden", true},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := isLongPressedName(tt.name, tt.typed); got != tt.want {
				t.Errorf("isLongPressedName() = %v, want %v", got, tt.want)
			}
		})
	}
}
