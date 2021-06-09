package main

import (
	"algorithm/stack"
	"fmt"
)

func main() {
	fmt.Println(stack.RemoveDuplicates(""))
	fmt.Println(stack.RemoveDuplicates("a"))
	fmt.Println(stack.RemoveDuplicates("ab"))
	fmt.Println(stack.RemoveDuplicates("abb"))
	fmt.Println(stack.RemoveDuplicates("abbb"))
	fmt.Println(stack.RemoveDuplicates("abbbb"))
	fmt.Println(stack.RemoveDuplicates("abcc"))
	fmt.Println(stack.RemoveDuplicates("abccd"))
	fmt.Println(stack.RemoveDuplicates("abccb"))
	fmt.Println(stack.RemoveDuplicates("abccbeebggbe"))
}
