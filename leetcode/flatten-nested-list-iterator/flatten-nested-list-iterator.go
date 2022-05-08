package main

// source: https://leetcode.com/problems/flatten-nested-list-iterator/

// NestedInteger This is the interface that allows for creating nested lists.
// You should not implement it, or speculate about its implementation
type NestedInteger struct {
	isInt      bool
	val        int
	nestedList []NestedInteger
}

// IsInteger Return true if this NestedInteger holds a single integer, rather than a nested root.
func (this NestedInteger) IsInteger() bool {
	return this.isInt
}

// GetInteger Return the single integer that this NestedInteger holds, if it holds a single integer
// The result is undefined if this NestedInteger holds a nested root
// So before calling this method, you should have a check
func (this NestedInteger) GetInteger() int {
	if this.isInt {
		return this.val
	}
	return -1
}

// SetInteger Set this NestedInteger to hold a single integer.
func (n NestedInteger) SetInteger(value int) {
	n.val = value
}

// Add Set this NestedInteger to hold a nested root and adds a nested integer to it.
func (this NestedInteger) Add(elem NestedInteger) {

}

// GetList Return the nested root that this NestedInteger holds, if it holds a nested root
// The root length is zero if this NestedInteger holds a single integer
// You can access NestedInteger's List element directly if you want to modify it
func (this NestedInteger) GetList() []*NestedInteger {
	return []*NestedInteger{}
}

// Actual solution is below
/*
type Stack struct {
	values []*NestedInteger
}

func NewStack() *Stack {
	return &Stack{}
}

func (s *Stack) Push(val *NestedInteger) {
	s.values = append(s.values, val)
}

func (s *Stack) Pop() (val *NestedInteger, ok bool) {
	if len(s.values) == 0 {
		return val, false
	}
	val = s.values[len(s.values)-1]
	s.values = s.values[:len(s.values)-1]
	return val, true
}

func (s *Stack) Len() int {
	return len(s.values)
}

func (s *Stack) Top() *NestedInteger {
	return s.values[len(s.values)-1]
}

type NestedIterator struct {
	stack *Stack
	root  []*NestedInteger
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	res := new(NestedIterator)
	res.root = nestedList
	res.stack = NewStack()
	res.stack.Push(nestedList[0])
	return res
}

func (it *NestedIterator) Next() int {
	var res int
	cur := it.stack.Pop()
	if cur.IsInteger() {
		res = cur.GetInteger()
		root := it.stack.Top()
	}
}

func (it *NestedIterator) HasNext() bool {

}
*/

// This solution requires O(n) memory and uses DFS when building iterator. Alternative way
// is to use stack to track current position relative to the root NestedInteger

type NestedIterator struct {
	values []int
	cur    int
}

func Constructor(nestedList []*NestedInteger) *NestedIterator {
	it := new(NestedIterator)
	var insert func(it *NestedIterator, val *NestedInteger)

	insert = func(it *NestedIterator, val *NestedInteger) {
		if val.IsInteger() {
			it.values = append(it.values, val.GetInteger())
			return
		}
		for _, i := range val.GetList() {
			insert(it, i)
		}
	}

	for _, i := range nestedList {
		insert(it, i)
	}

	it.cur = -1
	return it
}

func (it *NestedIterator) Next() int {
	return it.values[it.cur]
}

func (it *NestedIterator) HasNext() bool {
	if it.cur+1 < len(it.values) {
		return true
	}
	return false
}
