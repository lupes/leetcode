package question_201_210

import (
	"reflect"
)

func isIsomorphic(s string, t string) bool {
	var flags1 = make(map[int32][]int)
	var flags2 = make(map[int32][]int)
	for i, c := range s {
		flags1[c] = append(flags1[c], i)
	}
	for i, c := range t {
		flags2[c] = append(flags2[c], i)
	}
	for i, c := range s {
		if !reflect.DeepEqual(flags1[c], flags2[int32(t[i])]) {
			return false
		}
	}
	return true
}
