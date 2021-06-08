package main

import (
	"algorithm/llist"
)

func main() {
	y1 := llist.New([]int{0})
	y2 := llist.New([]int{})
	y3 := llist.MergeTwoLists(y1, y2)
	llist.Print(y3)
}
