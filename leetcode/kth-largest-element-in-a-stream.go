package main

import (
	"container/heap"
	"fmt"
)

type KthLargest struct {
	values []int
	k      int
}

func Constructor(k int, values []int) KthLargest {
	l := KthLargest{}
	l.values = append(make([]int, 0, len(values)), values...)
	l.k = k
	heap.Init(&l)
	for len(l.values) > k {
		heap.Pop(&l)
	}
	return l
}

func (this KthLargest) Len() int           { return len(this.values) }
func (this KthLargest) Less(i, j int) bool { return this.values[i] < this.values[j] }
func (this KthLargest) Swap(i, j int) {
	this.values[i], this.values[j] = this.values[j], this.values[i]
}

func (this *KthLargest) Push(x any) {
	this.values = append(this.values, x.(int))
}

func (this *KthLargest) Pop() any {
	old := *this
	n := len(old.values)
	x := old.values[n-1]
	this.values = old.values[0 : n-1]
	return x
}

func (this *KthLargest) Add(val int) int {
	heap.Push(this, val)
	for len(this.values) > this.k {
		heap.Pop(this)
	}
	return this.values[0]
}

func main() {
	var l3 = Constructor(2, []int{0})
	fmt.Println(l3.Add(-1))
	fmt.Println(l3.Add(1))
	fmt.Println(l3.Add(-2))
	fmt.Println(l3.Add(-4))
	fmt.Println(l3.Add(3))
	println()

	var l2 = Constructor(1, []int{})
	fmt.Println(l2.Add(-3))
	fmt.Println(l2.Add(-2))
	fmt.Println(l2.Add(-4))
	fmt.Println(l2.Add(0))
	fmt.Println(l2.Add(4))
	println()

	var l = Constructor(3, []int{4, 5, 8, 2})
	fmt.Println(l.Add(3))  // return 4
	fmt.Println(l.Add(5))  // return 5
	fmt.Println(l.Add(10)) // return 5
	fmt.Println(l.Add(9))  // return 8
	fmt.Println(l.Add(4))  // return 8
}
