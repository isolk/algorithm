package main

import (
	"algorithm/sort"
	"fmt"
)

func main() {
	y := []int{5, 100, 3, 2, 1}
	// y := []int{0, 1, 2, 3, 4}
	sort.SelectSort(y)
	fmt.Println(y)
}
