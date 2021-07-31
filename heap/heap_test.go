package heap

import (
	"testing"
)

func TestHeap(t *testing.T) {
	a := New(100)
	a.Insert(1)
	a.Top()
}
