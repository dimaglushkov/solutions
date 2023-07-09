package monoqueue

import "sort"

type MonoQueue struct {
	h *elem
	s int
}

type elem struct {
	val  int
	next *elem
}

func NewMonoQueue(values ...int) MonoQueue {
	if len(values) == 0 {
		return MonoQueue{}
	}
	//sort.Sort(sort.Reverse(sort.IntSlice(values)))
	sort.Ints(values)

	var prev, head *elem
	for _, i := range values {
		e := &elem{val: i, next: prev}
		prev = e
		head = e
	}

	return MonoQueue{
		h: head,
		s: len(values),
	}
}

func (mq *MonoQueue) Len() int {
	return mq.s
}

func (mq *MonoQueue) Push(x int) {
	if mq.h == nil {
		e := &elem{val: x}
		mq.h = e
	} else if x > mq.h.val {
		ne := &elem{val: x, next: mq.h}
		mq.h = ne
	} else {
		cur := mq.h
		for cur.next != nil && cur.next.val > x {
			cur = cur.next
		}
		ne := &elem{val: x, next: cur.next}
		cur.next = ne
	}
	mq.s++
}

func (mq *MonoQueue) Pop() int {
	x := mq.h.val
	mq.h = mq.h.next
	mq.s--
	return x
}

func (mq *MonoQueue) Peek() int {
	return mq.h.val
}
