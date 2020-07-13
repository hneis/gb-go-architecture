// Package main provides ...
package main

import (
	"fmt"

	"github.com/hneis/gb-go-architecture/lesson-5/homework/structures"
)

func main() {
	l := structures.List{}
	l.Append("Hello")
	l.Append("Hello1")
	l.Append("Hello2")

	l.Remove(l.Head)

	fmt.Println(l.Count)

	fmt.Println("Stack")
	stack := structures.Stack{}
	stack.Push("H")
	stack.Push("e")
	stack.Push("l")
	stack.Push("l")
	stack.Push("o")

	for pop := stack.Pop(); pop != nil; pop = stack.Pop() {
		fmt.Println(pop)
	}

	fmt.Println("Queue")
	queue := structures.Queue{}
	queue.Push("H")
	queue.Push("e")
	queue.Push("l")
	queue.Push("l")
	queue.Push("o")

	for pop := queue.Pop(); pop != nil; pop = queue.Pop() {
		fmt.Println(pop)
	}
}
