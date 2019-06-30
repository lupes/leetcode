package question_201_210

import "testing"

func TestTrie(t *testing.T) {
	trie := Constructor()
	trie.Insert("apple")
	if ok := trie.Search("apple"); !ok {
		t.Errorf("Search(apple) = %v, want %v", ok, true)
	}

	if ok := trie.Search("app"); ok {
		t.Errorf("Search(apple) = %v, want %v", ok, false)
	}

	if ok := trie.StartsWith("app"); !ok {
		t.Errorf("Search(apple) = %v, want %v", ok, true)
	}

	trie.Insert("app")
	if ok := trie.Search("app"); !ok {
		t.Errorf("Search(apple) = %v, want %v", ok, true)
	}

}
