package main

// source: https://leetcode.com/problems/peeking-iterator/

/*   Below is the interface for Iterator, which is already defined for you.
 *
 *   type Iterator struct {
 *
 *   }
 *
 *   func (this *Iterator) hasNext() bool {
 *		// Returns true if the iteration has more elements.
 *   }
 *
 *   func (this *Iterator) next() int {
 *		// Returns the next element in the iteration.
 *   }
 */

type Iterator struct{}

func (this *Iterator) hasNext() (res bool) {
	//stub
	return
}

func (this *Iterator) next() (res int) {
	//stub
	return
}

type PeekingIterator struct {
	i *Iterator
}

func Constructor(iter *Iterator) *PeekingIterator {
	return &PeekingIterator{i: iter}
}

func (this *PeekingIterator) hasNext() bool {
	return this.i.hasNext()
}

func (this *PeekingIterator) next() int {
	return this.i.next()
}

// creating a copy of the current iterator
// and returning result of .next() call on the copy
func (this *PeekingIterator) peek() int {
	tmp := *this.i
	return tmp.next()
}
