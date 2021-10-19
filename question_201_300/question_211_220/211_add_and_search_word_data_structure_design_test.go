package question_211_220

import (
	"fmt"
	"testing"
)

func TestWordDictionary(t *testing.T) {
	wd := Constructor()
	wd.AddWord("bad")
	wd.AddWord("dad")
	wd.AddWord("mad")
	fmt.Printf("%v\n", wd.Search("pad"))
	fmt.Printf("%v\n", wd.Search("bad"))
	fmt.Printf("%v\n", wd.Search(".ad"))
	fmt.Printf("%v\n", wd.Search("b.."))
}
