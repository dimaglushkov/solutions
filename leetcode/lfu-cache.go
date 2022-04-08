package main

// source: https://leetcode.com/problems/lfu-cache/

/*
// First working version. Obvious bottleneck on bumping node O(n) in worst case
import "container/list"

type LFUCacheNode struct {
	key, value, accesses int
}

type LFUCache struct {
	cap    int
	kv     map[int]*list.Element
	buckets *list.List
}

func NewLFUCache(capacity int) LFUCache {
	cache := LFUCache{
		cap:    capacity,
		kv:     make(map[int]*list.Element, capacity),
		buckets: list.New(),
	}
	return cache
}

func (this *LFUCache) Get(key int) int {
	if elem, ok := this.kv[key]; ok {
		this.bumpElem(elem)
		return elem.Value.(*LFUCacheNode).value
	}
	return -1
}

func (this *LFUCache) Put(key int, value int) {
	if this.cap == 0 {
		return
	}

	if elem, ok := this.kv[key]; ok {
		node := elem.Value.(*LFUCacheNode)
		node.value = value
		this.bumpElem(elem)
		return
	}

	if len(this.kv) == this.cap {
		tail := this.buckets.Back()
		delete(this.kv, tail.Value.(*LFUCacheNode).key)
		this.buckets.Remove(tail)
	}
	node := &LFUCacheNode{
		key:   key,
		value: value,
	}
	elem := this.buckets.PushBack(node)
	this.kv[key] = elem
	this.bumpElem(elem)
}

func (this *LFUCache) bumpElem(elem *list.Element) {
	node := elem.Value.(*LFUCacheNode)
	node.accesses++
	whereToMove := elem
	for whereToMove != this.buckets.Front() && whereToMove.Prev().Value.(*LFUCacheNode).accesses <= node.accesses {
		whereToMove = whereToMove.Prev()
	}
	this.buckets.MoveBefore(elem, whereToMove)
}
*/

import (
	"container/list"
)

type LFUCacheNode struct {
	key, value, accesses int
}

type LFUCache struct {
	cap, len, minBucketId int
	kv                    map[int]*list.Element
	buckets               map[int]*list.List
}

// NewLFUCache To run this code at leetcode replace the name of this method with Constructor
func NewLFUCache(capacity int) LFUCache {
	cache := LFUCache{
		cap:     capacity,
		kv:      make(map[int]*list.Element, capacity),
		buckets: make(map[int]*list.List),
	}
	cache.buckets[cache.minBucketId] = list.New()
	return cache
}

func (c *LFUCache) Get(key int) int {
	if elem, ok := c.kv[key]; ok {
		c.bumpElem(elem)
		return elem.Value.(*LFUCacheNode).value
	}
	return -1
}

func (c *LFUCache) Put(key int, value int) {
	if c.cap == 0 {
		return
	}

	if elem, ok := c.kv[key]; ok {
		node := elem.Value.(*LFUCacheNode)
		node.value = value
		c.bumpElem(elem)
		return
	}

	if c.len == c.cap {
		tail := c.buckets[c.minBucketId].Back()
		delete(c.kv, tail.Value.(*LFUCacheNode).key)
		c.buckets[c.minBucketId].Remove(tail)
		c.len--
	}

	c.minBucketId = 0
	elem := c.buckets[c.minBucketId].PushFront(&LFUCacheNode{
		key:   key,
		value: value,
	})
	c.kv[key] = elem
	c.len++
}

func (c *LFUCache) bumpElem(elem *list.Element) {
	node := elem.Value.(*LFUCacheNode)
	c.buckets[node.accesses].Remove(elem)
	if node.accesses == c.minBucketId && c.buckets[node.accesses].Len() == 0 {
		c.minBucketId++
	}

	node.accesses++

	if c.buckets[node.accesses] == nil {
		c.buckets[node.accesses] = list.New()
	}
	c.kv[node.key] = c.buckets[node.accesses].PushFront(node)

}

func main() {
	c4 := NewLFUCache(10)
	c4.Put(10, 13)
	c4.Put(3, 17)
	c4.Put(6, 11)
	c4.Put(10, 5)
	c4.Put(9, 10)
	println("1", c4.Get(13))
	c4.Put(2, 19)
	println("2", c4.Get(2))
	println("3", c4.Get(3))
	c4.Put(5, 25)
	println("4", c4.Get(8))
	c4.Put(9, 22)
	c4.Put(5, 5)
	c4.Put(1, 30)
	println("5", c4.Get(11))
	c4.Put(9, 12)
	println("6", c4.Get(7))
	println("7", c4.Get(5))
	println("8", c4.Get(8))
	println("9", c4.Get(9))
	c4.Put(4, 30)
	c4.Put(9, 3)
	println("10", c4.Get(9))
	println("11", c4.Get(10))
	println("12", c4.Get(10))
	c4.Put(6, 14)
	c4.Put(3, 1)
	println("13", c4.Get(3))
	c4.Put(10, 11)
	println("14", c4.Get(8))
	c4.Put(2, 14)
	println("15", c4.Get(1))
	println("16", c4.Get(5))
	println("17", c4.Get(4))
	c4.Put(11, 4)
	c4.Put(12, 24)
	c4.Put(5, 18)
	println("18", c4.Get(13))
	c4.Put(7, 23)
	println("19", c4.Get(8))
	println("20", c4.Get(12))
	c4.Put(3, 27)
	c4.Put(2, 12)
	println("21", c4.Get(5))
	c4.Put(2, 9)
	c4.Put(13, 4)
	c4.Put(8, 18)
	c4.Put(1, 7)
	println("22", c4.Get(6))
	c4.Put(9, 29)
	c4.Put(8, 21)
	println("23", c4.Get(5))
	c4.Put(6, 30)
	c4.Put(1, 12)
	println("24", c4.Get(10))
	c4.Put(4, 15)
	c4.Put(7, 22)
	c4.Put(11, 26)
	c4.Put(8, 17)
	c4.Put(9, 29)
	println("25", c4.Get(5))
	c4.Put(3, 4)
	c4.Put(11, 30)
	println("26", c4.Get(12))
	c4.Put(4, 29)
	println("27", c4.Get(3))
	println("28", c4.Get(9))
	println("29", c4.Get(6))
	c4.Put(3, 4)
	println("30", c4.Get(1))
	println("31", c4.Get(10))
	c4.Put(3, 29)
	c4.Put(10, 28)
	c4.Put(1, 20)
	c4.Put(11, 13)
	println("32", c4.Get(3))
	c4.Put(3, 12)
	c4.Put(3, 8)
	c4.Put(10, 9)
	c4.Put(3, 26)
	println("33", c4.Get(8))
	println("34", c4.Get(7))
	println("35", c4.Get(5))
	c4.Put(13, 17)
	c4.Put(2, 27)
	c4.Put(11, 15)
	println("36", c4.Get(12))
	c4.Put(9, 19)
	c4.Put(2, 15)
	c4.Put(3, 16)
	println("37", c4.Get(1))
	c4.Put(12, 17)
	c4.Put(9, 1)
	c4.Put(6, 19)
	println("38", c4.Get(4))
	println("39", c4.Get(5))
	println("40", c4.Get(5))
	c4.Put(8, 1)
	c4.Put(11, 7)
	c4.Put(5, 2)
	c4.Put(9, 28)
	println("41", c4.Get(1))
	c4.Put(2, 2)
	c4.Put(7, 4)
	c4.Put(4, 22)
	c4.Put(7, 24)
	c4.Put(9, 26)
	c4.Put(13, 28)
	c4.Put(11, 26)

	c3 := NewLFUCache(1)
	c3.Put(2, 1)
	println(c3.Get(2))
	c3.Put(3, 2)
	println(c3.Get(2))
	println(c3.Get(3))
	println()

	c2 := NewLFUCache(0)
	c2.Put(0, 0)
	println(c2.Get(0))
	println()

	c1 := NewLFUCache(2)
	c1.Put(1, 1)       // cache=[1,_], cnt(1)=1
	c1.Put(2, 2)       // cache=[2,1], cnt(2)=1, cnt(1)=1
	println(c1.Get(1)) // return 1
	c1.Put(3, 3)       // 2 is the c1 key because cnt(2)=1 is the smallest, invalidate 2.
	println(c1.Get(2)) // return -1 (not found)
	println(c1.Get(3)) // return 3
	c1.Put(4, 4)       // Both 1 and 3 have the same cnt, but 1 is LRU, invalidate 1.
	println(c1.Get(1)) // return -1 (not found)
	println(c1.Get(3)) // return 3
	println(c1.Get(4)) // return 4

}
