package main

import (
	"container/heap"
	"fmt"
)

type KthLargest struct {
	values []int
	k      int
}

// Constructor To run this code at leetcode replace the name of this method with Constructor
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

func (l KthLargest) Len() int           { return len(l.values) }
func (l KthLargest) Less(i, j int) bool { return l.values[i] < l.values[j] }
func (l KthLargest) Swap(i, j int) {
	l.values[i], l.values[j] = l.values[j], l.values[i]
}

func (l *KthLargest) Push(x any) {
	l.values = append(l.values, x.(int))
}

func (l *KthLargest) Pop() any {
	old := *l
	n := len(old.values)
	x := old.values[n-1]
	l.values = old.values[0 : n-1]
	return x
}

func (l *KthLargest) Add(val int) int {
	heap.Push(l, val)
	for len(l.values) > l.k {
		heap.Pop(l)
	}
	return l.values[0]
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
