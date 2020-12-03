package question_201_210

import (
	"testing"
)

func TestTrie(t *testing.T) {
	//trie := Constructor()
	//trie.Insert("apple")
	//
	//if ok := trie.Search("apple"); !ok {
	//	t.Errorf("Search(apple) = %v, want %v", ok, true)
	//}
	//
	//if ok := trie.Search("app"); ok {
	//	t.Errorf("Search(app) = %v, want %v", ok, false)
	//}
	//
	//if ok := trie.StartsWith("app"); !ok {
	//	t.Errorf("StartsWith(app) = %v, want %v", ok, true)
	//}
	////
	//trie.Insert("app")
	//if ok := trie.Search("app"); !ok {
	//	t.Errorf("Search(apple) = %v, want %v", ok, true)
	//}
	//
	//if ok, got := trie.Search("app"), true; ok != got {
	//	t.Errorf("Search(apple) = %v, want %v", ok, got)
	//}
	//
	//if ok, got := trie.Search("ab"), false; ok != got {
	//	t.Errorf("Search(apple) = %v, want %v", ok, got)
	//}
	//
	//if ok, got := trie.Search("appl"), false; ok != got {
	//	t.Errorf("Search(apple) = %v, want %v", ok, got)
	//}

	this := Constructor()
	this.Insert("app")
	this.Insert("apple")
	if ok, got := this.Search("app"), true; ok != got {
		t.Errorf("Search(apple) = %v, want %v", ok, got)
	}
}
