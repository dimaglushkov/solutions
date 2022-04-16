package main

import "fmt"

// source: https://leetcode.com/problems/lru-cache/

type LRUCacheNode struct {
	key, value int
	prev, next *LRUCacheNode
}

type LRUCache struct {
	cap        int
	values     map[int]*LRUCacheNode
	head, tail *LRUCacheNode
}

func Constructor(capacity int) LRUCache {
	cache := LRUCache{
		cap:    capacity,
		values: make(map[int]*LRUCacheNode, capacity),
		head:   new(LRUCacheNode),
		tail:   new(LRUCacheNode),
	}
	cache.head.next = cache.tail
	cache.tail.prev = cache.head
	return cache
}

func (c *LRUCache) Get(key int) int {
	if node, ok := c.values[key]; ok {
		c.bumpNode(node)
		return node.value
	}
	return -1
}

func (c *LRUCache) Put(key int, value int) {
	if node, ok := c.values[key]; ok {
		node.value = value
		c.bumpNode(node)
		return
	}

	node := &LRUCacheNode{
		key:   key,
		value: value,
	}
	c.values[key] = node
	c.bumpNode(node)
	if len(c.values) > c.cap {
		delete(c.values, c.tail.prev.key)
		c.tail.prev.prev.next = c.tail
		c.tail.prev = c.tail.prev.prev
	}
}

func (c *LRUCache) bumpNode(node *LRUCacheNode) {
	if c.head.next == node {
		return
	}

	if node.prev != nil {
		node.prev.next = node.next
	}
	if node.next != nil {
		node.next.prev = node.prev
	}

	node.prev = c.head
	node.next = c.head.next
	c.head.next.prev = node
	c.head.next = node
}

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
