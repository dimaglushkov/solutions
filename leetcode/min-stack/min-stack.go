package main

import "math"

// source: https://leetcode.com/problems/min-stack/

type MinStack struct {
	data []int
	mins []int
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func Constructor() MinStack {
	ms := MinStack{
		data: []int{},
		mins: []int{math.MaxInt},
	}
	return ms
}

func (ms *MinStack) Push(val int) {
	ms.data = append(ms.data, val)
	ms.mins = append(ms.mins, min(val, ms.mins[len(ms.data)-1]))
}

func (ms *MinStack) Pop() {
	ms.data = ms.data[:len(ms.data)-1]
	ms.mins = ms.mins[:len(ms.mins)-1]
}

func (ms *MinStack) Top() int {
	return ms.data[len(ms.data)-1]
}

func (ms *MinStack) GetMin() int {
	return ms.mins[len(ms.data)]
}

/**
 * Your MinStack object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Push(val);
 * obj.Pop();
 * param_3 := obj.Top();
 * param_4 := obj.GetMin();
 */

func main() {
	ms := Constructor()
	ms.Push(4)
	ms.Push(3)
	ms.Push(8)
	ms.Push(6)
	ms.Push(1)
	ms.Push(2)
	return
}
