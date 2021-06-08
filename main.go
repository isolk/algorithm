package main

import (
	"algorithm/llist"
	"fmt"
)

func main() {
	y := llist.NewCycle([]int{}, -1)
	fmt.Println(llist.HasCycle(y))
}
