package main

// source: https://leetcode.com/problems/design-hashset/

// Ways to improve:
// 1. Use better collision resolving strategies
// 2. Rebuild hash-table when bucket size gets bigger than some value

const capacity = 100

type BucketNode struct {
	key  int
	next *BucketNode
}

type MyHashSet struct {
	buckets []*BucketNode
}

func Constructor() MyHashSet {
	set := MyHashSet{
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

func hashFunc(val int) (hash int) {
	return val % capacity
}

func (s *MyHashSet) Add(key int) {
	prevNode, hash := s.findPrev(key)
	if prevNode == nil || prevNode.key != -1 {
		return
	}
	node := &BucketNode{
		key:  key,
		next: s.buckets[hash],
	}
	s.buckets[hash] = node
}

func (s *MyHashSet) findPrev(key int) (*BucketNode, int) {
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

func (s *MyHashSet) Remove(key int) {
	node, hash := s.findPrev(key)
	if node == nil {
		s.buckets[hash] = s.buckets[hash].next
	} else if node.key != -1 {
		node.next = node.next.next
	}
}

func (s *MyHashSet) Contains(key int) bool {
	if tmp, _ := s.findPrev(key); tmp != nil && tmp.key == -1 {
		return false
	}
	return true
}

/**
 * Your s object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Add(key);
 * obj.Remove(key);
 * param_3 := obj.Contains(key);
 */
func main() {
	s := Constructor()
	s.Add(1)               // set = [1]
	s.Add(2)               // set = [1, 2]
	println(s.Contains(1)) // return True
	println(s.Contains(3)) // return False, (not found)
	s.Add(2)               // set = [1, 2]
	println(s.Contains(2)) // return True
	s.Remove(2)            // set = [1]
	println(s.Contains(2)) // return False, (already removed)
}
