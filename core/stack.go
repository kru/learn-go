package core

import (
	"fmt"
	"log"
)

type NodeStack struct {
	value string
	prev  *NodeStack
	head  *NodeStack
}

func (n *NodeStack) Push(data string) {
	node := &NodeStack{value: data}
	fmt.Println("adding data", data)
	if n.head == nil {
		n.head = node
		n = node
	} else {
		node.prev = n.head
		n.head = node

	}
}

func (n *NodeStack) Pop() {

	if n.head == nil {
		log.Fatalf("there is no stack yet: %v\n", n)
	} else {
		n.head = n.head.prev
	}
}

func (n *NodeStack) Peek() {

	fmt.Println()
	if n.head == nil {
		fmt.Printf("there is no data in stack: %+v\n", n)
	} else {
		tmp := n.head
		size := 0
		for tmp.prev != nil {
			fmt.Printf("data for node: %s\n", tmp.value)
			tmp = tmp.prev
			size++
		}
		fmt.Printf("data for node: %s\n", tmp.value)
		fmt.Printf("stack size: %d\n", size+1)
	}
}

func InitStack() *NodeStack {
	return &NodeStack{}
}

func NewStack() {
	st := InitStack()

	st.Push("A")
	st.Peek()
	st.Push("B")
	st.Push("C")
	st.Push("D")
	st.Peek()
	st.Pop()
	st.Peek()
	st.Pop()
	st.Peek()

}
