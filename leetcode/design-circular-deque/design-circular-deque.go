package main

// source: https://leetcode.com/problems/design-circular-deque/
type MyCircularDeque struct {
	val []int
	n   int
}

func Constructor(k int) MyCircularDeque {
	return MyCircularDeque{val: make([]int, 0, k), n: k}
}

func (d *MyCircularDeque) InsertFront(value int) bool {
	if d.IsFull() {
		return false
	}

	d.val = append([]int{value}, d.val...)
	return true
}

func (d *MyCircularDeque) InsertLast(value int) bool {
	if d.IsFull() {
		return false
	}

	d.val = append(d.val, value)
	return true
}

func (d *MyCircularDeque) DeleteFront() bool {
	if d.IsEmpty() {
		return false
	}
	d.val = d.val[1:]
	return true
}

func (d *MyCircularDeque) DeleteLast() bool {
	if d.IsEmpty() {
		return false
	}
	d.val = d.val[:len(d.val)-1]
	return true
}

func (d *MyCircularDeque) GetFront() int {
	if d.IsEmpty() {
		return -1
	}
	return d.val[0]
}

func (d *MyCircularDeque) GetRear() int {
	if d.IsEmpty() {
		return -1
	}
	return d.val[len(d.val)-1]
}

func (d *MyCircularDeque) IsEmpty() bool {
	return len(d.val) == 0
}

func (d *MyCircularDeque) IsFull() bool {
	return len(d.val) == d.n
}

/**
 * Your MyCircularDeque object will be instantiated and called as such:
 * obj := Constructor(k);
 * param_1 := obj.InsertFront(value);
 * param_2 := obj.InsertLast(value);
 * param_3 := obj.DeleteFront();
 * param_4 := obj.DeleteLast();
 * param_5 := obj.GetFront();
 * param_6 := obj.GetRear();
 * param_7 := obj.IsEmpty();
 * param_8 := obj.IsFull();
 */
