package main

import (
	"math/rand"
	"time"
)

// source: https://leetcode.com/problems/insert-delete-getrandom-o1/

type RandomizedSet struct {
	data map[int]bool
	keys []int
}

func Constructor() RandomizedSet {
	rand.Seed(time.Now().Unix())
	return RandomizedSet{
		data: make(map[int]bool),
		keys: make([]int, 0),
	}
}

func (rs *RandomizedSet) Insert(val int) bool {
	if !rs.data[val] {
		rs.data[val] = true
		rs.keys = append(rs.keys, val)
		return true
	}
	return false
}

func (rs *RandomizedSet) Remove(val int) bool {
	if rs.data[val] {
		delete(rs.data, val)
		return true
	}
	return false
}

func (rs *RandomizedSet) GetRandom() int {
	for len(rs.keys) > 0 {
		i := rand.Intn(len(rs.keys))
		if v := rs.keys[i]; rs.data[v] {
			return v
		}
		rs.keys = append(rs.keys[:i], rs.keys[i+1:]...)
	}
	return -1
}

func main() {
	rs := Constructor()
	println(rs.Insert(1))
	println(rs.Remove(2))
	println(rs.Insert(2))
	println(rs.GetRandom())
	println(rs.Remove(1))
	println(rs.Insert(2))
	println(rs.GetRandom())

}
