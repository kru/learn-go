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
	return
}

func (d *DoublyLinkedList) AddBackNode(data string) {
	node := &Node{data: data}

	if d.head == nil {
		d.head = node
		d.tail = node
	} else {
		currentNode := d.head
		for currentNode.next != nil {
			currentNode = currentNode.next
		}
		node.prev = d.tail
		d.tail.next = node
		d.tail = node
	}
	d.ln++
	return
}

func (d *DoublyLinkedList) Size() int {
	return d.ln
}

func (d *DoublyLinkedList) TraverseForward() error {
	if d.head == nil {
		return fmt.Errorf("Error no double linked list data %v\n", d.head)
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

func (d *DoublyLinkedList) Reverse() error {
	if d.tail == nil {
		return fmt.Errorf("can not reverse DLL, no data")
	}
	var clone = &DoublyLinkedList{}

	temp := d.tail
	fmt.Printf("TEMP: %v\n", temp.prev)
	for temp.prev != nil {
		if clone.head == nil {
			clone.head = temp
			clone.tail = temp
		} else {
			temp.prev = clone.tail
			clone.tail.next = temp
		}
		temp = temp.prev
	}
	return nil
}
