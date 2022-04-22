package main

// source: https://leetcode.com/problems/design-hashset/

// Ways to improve:
// 1. Use better collision resolving strategies
// 2. Rebuild hash-table when bucket size gets bigger than some value

const capacity = 10000

type BucketNode struct {
	key, value int
	next       *BucketNode
}

type MyHashMap struct {
	buckets []*BucketNode
}

func Constructor() MyHashMap {
	set := MyHashMap{
		buckets: make([]*BucketNode, capacity),
	}
	lastNode := &BucketNode{
		next: nil,
		key:  -1,
	}
	for i := 0; i < capacity; i++ {
		set.buckets[i] = lastNode
	}
	return set
}

func hashFunc(key int) (hash int) {
	return key % capacity
}

func (s *MyHashMap) findPrev(key int) (*BucketNode, int) {
	hash := hashFunc(key)
	prevNode := s.buckets[hash]
	if prevNode.key == key {
		return nil, hash
	}
	for prevNode.next != nil && prevNode.next.key != key {
		prevNode = prevNode.next
	}
	return prevNode, hash
}

func (s *MyHashMap) Put(key, value int) {
	prevNode, hash := s.findPrev(key)
	if prevNode == nil {
		s.buckets[hash].value = value
		return
	}
	if prevNode.key != -1 {
		prevNode.value = value
		return
	}
	node := &BucketNode{
		key:   key,
		value: value,
		next:  s.buckets[hash],
	}
	s.buckets[hash] = node
}

func (s *MyHashMap) Get(key int) int {
	var prev *BucketNode
	var hash int
	if prev, hash = s.findPrev(key); prev != nil && prev.key == -1 {
		return -1
	}
	if prev == nil {
		return s.buckets[hash].value
	}
	return prev.next.value
}

func (s *MyHashMap) Remove(key int) {
	node, hash := s.findPrev(key)
	if node == nil {
		s.buckets[hash] = s.buckets[hash].next
	} else if node.key != -1 {
		node.next = node.next.next
	}
}

func main() {
	m := Constructor()
	m.Put(1, 1)       // The map is now [[1,1]]
	m.Put(2, 2)       // The map is now [[1,1], [2,2]]
	println(m.Get(1)) // return 1, The map is now [[1,1], [2,2]]
	println(m.Get(3)) // return -1 (i.e., not found), The map is now [[1,1], [2,2]]
	m.Put(2, 1)       // The map is now [[1,1], [2,1]] (i.e., update the existing value)
	println(m.Get(2)) // return 1, The map is now [[1,1], [2,1]]
	m.Remove(2)       // remove the mapping for 2, The map is now [[1,1]]
	println(m.Get(2)) // return -1 (i.e., not found), The map is now [[1,1]]
}
