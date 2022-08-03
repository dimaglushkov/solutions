package main

// source: https://leetcode.com/problems/my-calendar-i/

type Node struct {
	start, end  int
	left, right *Node
}

func insert(head **Node, start, end int) bool {
	switch {
	case *head == nil:
		*head = &Node{start: start, end: end}
		return true
	case end <= (*head).start:
		return insert(&((*head).left), start, end)
	case start >= (*head).end:
		return insert(&((*head).right), start, end)
	default:
		return false
	}
}

type MyCalendar struct {
	head *Node
}

func Constructor() MyCalendar {
	return MyCalendar{}
}

func (c *MyCalendar) Book(start int, end int) bool {
	return insert(&c.head, start, end)
}

/**
 * Your MyCalendar object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
