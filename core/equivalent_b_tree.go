package core

import (
	"fmt"
	"math/rand"
	"sort"
	"time"
)

type Tree struct {
	Left  *Tree
	Value int32
	Right *Tree
}

func New(k int32) *Tree {
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

func Walk(t *Tree, ch chan int32) {
	if t == nil {
		fmt.Println("tree is nil")
		return
	}
	fmt.Println(t)

	ch <- t.Value

	tl := t.Left
	for tl != nil {
		time.Sleep(100 * time.Millisecond)
		ch <- tl.Value
		tlr := tl.Right
		for tlr != nil {
			ch <- tlr.Value
			tlrl := tlr.Left
			for tlrl != nil {
				ch <- tlrl.Value
				tlrl = tlrl.Left
			}
			tlr = tlr.Right
		}
		tl = tl.Left
	}

	tr := t.Right
	for tr != nil {
		time.Sleep(100 * time.Millisecond)
		ch <- tr.Value
		trl := tr.Left
		for trl != nil {
			ch <- trl.Value
			trlr := trl.Right
			for trlr != nil {
				ch <- trlr.Value
				trlr = trlr.Right
			}
			trl = trl.Left
		}

		tr = tr.Right
	}

	close(ch)
}

func Same(t1, t2 *Tree) bool {
	var t1Vals, t2Vals foo
	ch1 := make(chan int32)
	ch2 := make(chan int32)

	go Walk(New(8), ch1)
	go Walk(New(8), ch2)

	for val := range ch1 {
		fmt.Println("t1 Val", val)
		t1Vals = append(t1Vals, val)
	}
	for v := range ch2 {
		t2Vals = append(t2Vals, v)
	}
	sort.Sort(t1Vals)
	sort.Sort(t2Vals)
	fmt.Println("t1Vals", t1Vals)
	fmt.Println("t2Vals", t2Vals)
	for idx := range t1Vals {
		if t1Vals[idx] != t2Vals[idx] {
			return false
		}
	}

	return true
}
