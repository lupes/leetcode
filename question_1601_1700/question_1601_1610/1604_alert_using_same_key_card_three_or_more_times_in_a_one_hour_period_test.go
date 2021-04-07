package question_1601_1610

import (
	"reflect"
	"testing"
)

func Test_alertNames(t *testing.T) {
	tests := []struct {
		keyName []string
		keyTime []string
		want    []string
	}{
		{
			[]string{"daniel", "daniel", "daniel", "luis", "luis", "luis", "luis"},
			[]string{"10:00", "10:40", "11:00", "09:00", "11:00", "13:00", "15:00"},
			[]string{"daniel"},
		},
		{
			[]string{"alice", "alice", "alice", "bob", "bob", "bob", "bob"},
			[]string{"12:01", "12:00", "18:00", "21:00", "21:20", "21:30", "23:00"},
			[]string{"bob"},
		},
		{
			[]string{"john", "john", "john"},
			[]string{"23:58", "23:59", "00:01"},
			nil,
		},
		{
			[]string{"leslie", "leslie", "leslie", "clare", "clare", "clare", "clare"},
			[]string{"13:00", "13:20", "14:00", "18:00", "18:51", "19:30", "19:49"},
			[]string{"clare", "leslie"},
		},
	}
	for _, tt := range tests {
		t.Run("test", func(t *testing.T) {
			if got := alertNames(tt.keyName, tt.keyTime); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("alertNames() = %v, want %v", got, tt.want)
			}
		})
	}
}
