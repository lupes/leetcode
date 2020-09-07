package question_1261_1270

import (
	"testing"

	. "github.com/lupes/leetcode/common"
)

func TestFindElements_Find(t *testing.T) {
	this := Constructor(NewNodeV2(-1, Null, -1))
	if got := this.Find(1); got != false {
		t.Errorf("Find() = %v", got)
	}

	if got := this.Find(2); got != true {
		t.Errorf("Find() = %v", got)
	}

	this = Constructor(NewNodeV2(-1, -1, -1, -1, -1))
	if got := this.Find(1); got != true {
		t.Errorf("Find() = %v", got)
	}

	if got := this.Find(3); got != true {
		t.Errorf("Find() = %v", got)
	}

	if got := this.Find(5); got != false {
		t.Errorf("Find() = %v", got)
	}
}
