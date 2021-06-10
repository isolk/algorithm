package main

import (
	"algorithm/queue"
	"fmt"
)

func main() {
	f := queue.NewQueue(3)
	f.Enqueue(1)
	f.Enqueue(2)
	f.Enqueue(3)
	fmt.Println(f.Dequeue())
	fmt.Println(f.Dequeue())
	f.Enqueue(4)
	f.Enqueue(5)
	fmt.Println(f.Dequeue())
	fmt.Println(f.Dequeue())
	fmt.Println(f.Dequeue())

	f.Enqueue(1)
	f.Enqueue(2)
	f.Enqueue(3)
	fmt.Println(f.Dequeue())
	fmt.Println(f.Dequeue())
	fmt.Println(f.Dequeue())
}
