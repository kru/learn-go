package core

// When they flip page 1, they see pages 2 and 3.
// Each page except the last page will always be printed on both sides.
// The last page may only be printed on the front, given the length of the book.
// If the book is  pages long, and a student wants to turn to page , what is the minimum number of pages to turn?
// They can start at the beginning or the end of the book.
// Given n and p, find and print the minimum number of pages that must be turned in order to arrive at page .
// Example n = 5 p = 3
func PageCount(n int32, p int32) int32 {

	if n <= 1 || p <= 1 {
		return 0
	}

	if (n+1)%2 != 0 {
		n = n + 1
	}
	front := p / 2
	back := ((n) - p) / 2

	if front > back {
		return back
	}
	return front

}
