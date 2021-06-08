package main

import (
	"algorithm/llist"
	"fmt"
)

func main() {
	y := llist.NewCycle([]int{1, 2, 3, 4, 5}, -1)
	fmt.Println(llist.HasCycle(y))
}
