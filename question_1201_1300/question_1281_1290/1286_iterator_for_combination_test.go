package question_1281_1290

import (
	"testing"
)

func TestCombinationIterator(t *testing.T) {
	iter := Constructor("abcde", 3)

	if res, want := iter.Next(), "abc"; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.HasNext(), true; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.Next(), "abd"; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.HasNext(), true; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.Next(), "abe"; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.HasNext(), true; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.Next(), "acd"; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.HasNext(), true; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.Next(), "ace"; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.HasNext(), true; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.Next(), "ade"; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.HasNext(), true; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.Next(), "bcd"; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.HasNext(), true; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.Next(), "bce"; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.HasNext(), true; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.Next(), "bde"; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.HasNext(), true; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.Next(), "cde"; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}

	if res, want := iter.HasNext(), false; res != want {
		t.Errorf("HasNext() = %v, want %v", res, want)
	}
}
