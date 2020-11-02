package question_1331_1340

import (
	"testing"
)

func TestConstructor(t *testing.T) {
	this := Constructor()
	this.Add(3)
	this.Add(0)
	this.Add(2)
	this.Add(5)
	this.Add(4)
	t.Logf("%d\n", this.GetProduct(2))
	t.Logf("%d\n", this.GetProduct(3))
	t.Logf("%d\n", this.GetProduct(4))
	this.Add(8)
	t.Logf("%d\n", this.GetProduct(2))
	this.Add(0)
	t.Logf("%d\n", this.GetProduct(1))
}
