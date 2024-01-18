package core

import (
	"fmt"
)

type node2 struct {
	prev *node2
	data string
	next *node2
}

type linkedList struct {
	length int
	head   *node2
	tail   *node2
}

func (ll *linkedList) toFront(data string) {
	n := &node2{data: data}
	if ll.head == nil {
		ll.head = n
		ll.tail = n
	} else {
		n.next = ll.head
		ll.head.prev = n
		ll.head = n
	}
	ll.length++
}

func (ll *linkedList) toBack(data string) {
	n := &node2{data: data}
	if ll.head == nil {
		ll.head = nil
		ll.tail = nil
	} else {
		n.prev = ll.tail
		ll.tail.next = n
		ll.tail = n
	}
	ll.length++
}

func (ll *linkedList) removeAt(idx int) {
	if idx > ll.length || idx < 0 {
		fmt.Println("out of range")
		return
	}

	// store current node to temporary node
	// remove node at idx
	// connect prev and next of deleted node
	tmp := ll.head
	var i int
	for tmp.next != nil {
		if i == idx {
			fmt.Println("removed data", tmp.data)
			nx := tmp.next
			tmp.prev.next = nx
			nx.prev = tmp.prev
			break
		}
		tmp = tmp.next
		i++
	}
	ll.length--
}

func (ll *linkedList) inserAt(idx int, data string) {
	if idx < 0 {
		fmt.Println("can not insert below 0")
		return
	}
	if idx == ll.length+1 {
		ll.toBack(data)
		return
	}
	if idx > ll.length+1 {
		fmt.Printf("the idx %d way higher than existing %d\n", idx, ll.length)
		return
	}

	tmp := ll.head
	i := 0

	node := &node2{data: data}
	for tmp.next != nil {
		if i == idx {
			fmt.Println("adding data", data)
			fmt.Println("current data", tmp.data)
			node.next = tmp
			node.prev = tmp.prev
			tmp.prev.next = node
			tmp.prev = node
		}
		tmp = tmp.next
		i++
	}
	ll.length++
}

func (ll *linkedList) get(data string) string {
	tmp := ll.head
	if tmp.data == data {
		return tmp.data
	}

	tmp = tmp.next
	var ln int
	for ln < ll.length && tmp != nil {
		if tmp.data == data {
			return tmp.data
		}
		tmp = tmp.next
		ln++
	}

	return fmt.Sprintf("your data %s not found\n", data)
}

func walkToTail(nd node2) {
	if nd.next == nil {
		fmt.Println(nd.data)
		return
	}
	fmt.Println(nd.data)
	walkToTail(*nd.next)
}

func walkToHead(nd node2) {
	if nd.prev == nil {
		fmt.Println(nd.data)
		return
	}
	fmt.Println(nd.data)
	walkToHead(*nd.prev)
}

func InitDoubly() {
	doubly := &linkedList{}

	doubly.toFront("C")
	doubly.toFront("B")

	doubly.toBack("D")
	doubly.toBack("E")
	doubly.toBack("F")
	fmt.Printf("ll : %+v\n", doubly.head)

	doubly.removeAt(2)
	walkToTail(*doubly.head)

	doubly.inserAt(2, "D")
	walkToTail(*doubly.head)

	fmt.Println("walk to tail")
	fmt.Printf("%+v\n", doubly.tail.prev)
	walkToHead(*doubly.tail)

	fmt.Println("getting data")
	dt := doubly.get("d")
	fmt.Printf("result: %s\n", dt)
}
