package main

import (
	"fmt"
	"github.com/kru/learn-go/core"
)

func main() {

	dll := core.InitDoublyLinkedList()

	fmt.Printf("Size of doubly linked ist: %d\n", dll.Size())

	fmt.Printf("Add Front Node: C\n")
	dll.AddFrontNode("C")

	fmt.Printf("Add End Node: D\n")
	dll.AddBackNode("D")

	fmt.Printf("Add Front Node: B\n")
	dll.AddFrontNode("B")

	fmt.Printf("Add End Node: E\n")
	dll.AddBackNode("E")

	fmt.Printf("Add Front Node: A\n")
	dll.AddFrontNode("A")

	fmt.Printf("Add End Node: F\n")
	dll.AddBackNode("F")

	fmt.Printf("Add End Node: G\n")
	dll.AddBackNode("G")

	fmt.Printf("Size of doubly linked ist: %d\n", dll.Size())

	dll.TraverseForward()
	dll.TraverseReverse()

	fmt.Println("Reversing DLL")
	dll.Reverse()

	dll.TraverseForward()
}
