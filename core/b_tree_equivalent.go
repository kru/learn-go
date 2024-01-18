package core

import (
	"fmt"
	"math/rand"
)

type Tree struct {
	Left  *Tree
	Value int32
	Right *Tree
}

func NewBTree(k int32) *Tree {
	var t *Tree
	for _, v := range rand.Perm(10) {
		t = insert(t, int32(1+v)*k)
	}
	return t
}

func insert(t *Tree, v int32) *Tree {
	if t == nil {
		return &Tree{nil, v, nil}
	}
	if v < t.Value {
		t.Left = insert(t.Left, v)
	} else {
		t.Right = insert(t.Right, v)
	}

	return t
}

func (t *Tree) String() string {
	if t == nil {
		return "()"
	}
	s := ""
	if t.Left != nil {
		s += t.Left.String() + " "
	}
	s += fmt.Sprint(t.Value)
	if t.Right != nil {
		s += " " + t.Right.String()
	}
	return "(" + s + ")"
}

func (t *Tree) Traverse() {
	if t.Left != nil {
		fmt.Println(t.Left.Value)
		t.Left.Traverse()
	}

	if t.Right != nil {
		fmt.Println(t.Right.Value)
		t.Right.Traverse()
	}
}

func (t *Tree) InOrder(tree *Tree) {
	if tree == nil {
		return
	} else {
		t.InOrder(tree.Left)
		fmt.Print(tree.Value, " ")
		t.InOrder(tree.Right)
	}
}

// t := core.NewBTree(8)
// ch := make(chan int32)
//
//	go func() {
//		core.WalkBTree(t, ch)
//		close(ch)
//	}()
//
//	for val := range ch {
//		fmt.Println(val)
//	}
func WalkBTree(t *Tree, ch chan int32) {
	if t == nil {
		return
	} else {
		WalkBTree(t.Left, ch)
		ch <- t.Value
		WalkBTree(t.Right, ch)
	}
}

func SameBTree(t1, t2 *Tree) bool {
	var t1Vals, t2Vals []int32
	ch1 := make(chan int32)
	ch2 := make(chan int32)
	go func() {
		WalkBTree(t1, ch1)
		close(ch1)
	}()

	go func() {
		WalkBTree(t2, ch2)
		close(ch2)
	}()

	for v := range ch1 {
		t1Vals = append(t1Vals, v)
	}

	for x := range ch2 {
		t2Vals = append(t2Vals, x)
	}

	for z := range t1Vals {
		if t1Vals[z] != t2Vals[z] {
			return false
		}
	}
	return true
}
