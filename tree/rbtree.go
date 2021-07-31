package tree

var Red = true
var Black = false

type RBNode struct {
	Key   int
	Val   interface{}
	Color bool
	Left  *RBNode
	Right *RBNode
}

type RBTree struct {
	Root *RBNode
}

func (r *RBTree) Find()
