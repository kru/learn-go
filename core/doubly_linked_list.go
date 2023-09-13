package core

import "fmt"

type Node struct {
	data string
	prev *Node
	next *Node
}

type DoublyLinkedList struct {
	ln   int
	head *Node
	tail *Node
}

func InitDoublyLinkedList() *DoublyLinkedList {
	return &DoublyLinkedList{}
}

func (d *DoublyLinkedList) AddFrontNode(data string) {
	node := &Node{data: data}

	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		node.next = d.head
		d.head.prev = node
		d.head = node
	}
	d.ln++
}

func (d *DoublyLinkedList) AddBackNode(data string) {
	node := &Node{data: data}

	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		node.prev = d.tail
		d.tail.next = node
		d.tail = node
	}
	d.ln++
}

func (d *DoublyLinkedList) Size() int {
	return d.ln
}

func (d *DoublyLinkedList) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("eror no double linked list data %v", d.head)
	}

	fmt.Println("TraverseForward")
	temp := d.head
	for temp.next != nil {
		fmt.Printf("Data: %s; previous %v; next %v\n", temp.data, temp.prev, temp.next)
		temp = temp.next
	}

	return nil
}

func (d *DoublyLinkedList) TraverseReverse() error {
	if d.tail == nil {
		return fmt.Errorf("error no double linked list tail")
	}

	fmt.Println("TraverseReverse")
	temp := d.tail
	for temp.prev != nil {
		fmt.Printf("Data: %s; previous %v; next %v\n", temp.data, temp.prev, temp.next)
		temp = temp.prev
	}

	return nil
}

func ReverseDll(d *DoublyLinkedList) (*DoublyLinkedList, error) {
	if d.tail == nil {
		return nil, fmt.Errorf("can not reverse DLL, no data")
	}
	var clone = &DoublyLinkedList{}

	temp := d.tail

	for temp.prev != nil {
		nd := &Node{data: temp.data}
		if clone.head == nil {
			clone.head = nd
			clone.tail = nd
			fmt.Printf("FIRST %+v\n", clone.head)
		} else {
			nd.prev = clone.tail
			clone.tail.next = nd
			clone.tail = nd
		}
		clone.ln++

		temp = temp.prev
	}

	if temp.prev == nil {
		nd := &Node{data: temp.data}
		nd.prev = clone.tail
		clone.tail.next = nil
		clone.tail = nd
		fmt.Printf("temp: %+v\n", clone.tail)
	}

	return clone, nil
}

func InitDLL() {
	dll := InitDoublyLinkedList()

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
	clone, err := ReverseDll(dll)

	if err != nil {
		fmt.Println("Reversing Error")
	}

	clone.TraverseForward()
	clone.TraverseReverse()
}
