package main

import "fmt"

// source: https://leetcode.com/problems/lru-cache/

type Node struct {
	key, value int
	prev, next *Node
}

type LRUCache struct {
	cap        int
	values     map[int]*Node
	head, tail *Node
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:    capacity,
		values: make(map[int]*Node, capacity),
		head:   new(Node),
		tail:   new(Node),
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (this *LRUCache) Get(key int) int {
	if node, ok := this.values[key]; ok {
		this.bumpNode(node)
		return node.value
	}
	return -1
}

func (this *LRUCache) Put(key int, value int) {
	if node, ok := this.values[key]; ok {
		node.value = value
		this.bumpNode(node)
		return
	}

	node := &Node{
		key:   key,
		value: value,
	}
	this.values[key] = node
	this.bumpNode(node)
	if len(this.values) > this.cap {
		delete(this.values, this.tail.prev.key)
		this.tail.prev.prev.next = this.tail
		this.tail.prev = this.tail.prev.prev
	}
}

func (this *LRUCache) bumpNode(node *Node) {
	if this.head.next == node {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	node.prev = this.head
	node.next = this.head.next
	this.head.next.prev = node
	this.head.next = node
}

/**
 * Your LRUCache object will be instantiated and called as such:
 * obj := Constructor(capacity);
 * param_1 := obj.Get(key);
 * obj.Put(key,value);
 */
func main() {
	var c1 = Constructor(2)
	c1.Put(1, 1)           // cache is {1=1}
	c1.Put(2, 2)           // cache is {1=1, 2=2}
	fmt.Println(c1.Get(1)) // return 1
	c1.Put(3, 3)           // LRU key was 2, evicts key 2, cache is {1=1, 3=3}
	fmt.Println(c1.Get(2)) // returns -1 (not found)
	c1.Put(4, 4)           // LRU key was 1, evicts key 1, cache is {4=4, 3=3}
	fmt.Println(c1.Get(1)) // return -1 (not found)
	fmt.Println(c1.Get(3)) // return 3
	fmt.Println(c1.Get(4)) // return 4

	println()

	//["LRUCache","put","put","put","put","get","get","get","get","put","get","get","get","get","get"]
	//[[3],		  [1,1],[2,2],[3,3],[4,4],[4],   [3],  [2],  [1],  [5,5],[1],  [2],  [3],  [4],  [5]]
	var c2 = Constructor(3)
	c2.Put(1, 1)
	c2.Put(2, 2)
	c2.Put(3, 3)
	c2.Put(4, 4)
	fmt.Println(c2.Get(4))
	fmt.Println(c2.Get(3))
	fmt.Println(c2.Get(2))
	fmt.Println(c2.Get(1))
	c2.Put(5, 5)
	fmt.Println(c2.Get(1))
	fmt.Println(c2.Get(2))
	fmt.Println(c2.Get(3))
	fmt.Println(c2.Get(4))
	fmt.Println(c2.Get(5))

}
