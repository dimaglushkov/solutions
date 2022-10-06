package main

import (
	"sort"
)

// source: https://leetcode.com/problems/time-based-key-value-store/

type event struct {
	ts  int
	val string
}

type TimeMap struct {
	m map[string][]event
}

func Constructor() TimeMap {
	return TimeMap{
		m: make(map[string][]event),
	}
}

func (this *TimeMap) Set(key string, value string, timestamp int) {
	if _, ok := this.m[key]; !ok {
		this.m[key] = make([]event, 0)
	}
	this.m[key] = append(this.m[key], event{ts: timestamp, val: value})
}

func (this *TimeMap) Get(key string, timestamp int) string {
	n := len(this.m[key])
	if n == 0 || timestamp < this.m[key][0].ts {
		return ""
	}

	i := sort.Search(n, func(i int) bool {
		return this.m[key][i].ts >= timestamp
	})

	if i < n && this.m[key][i].ts == timestamp {
		return this.m[key][i].val
	}

	return this.m[key][i-1].val
}

/**
 * Your TimeMap object will be instantiated and called as such:
 * obj := Constructor();
 * obj.Set(key,value,timestamp);
 * param_2 := obj.Get(key,timestamp);
 */

func main() {
	tm := Constructor()
	tm.Set("love", "high", 10)
	tm.Set("love", "low", 20)
	println(tm.Get("love", 5))
	println(tm.Get("love", 10))
	println(tm.Get("love", 15))
	println(tm.Get("love", 20))
	println(tm.Get("love", 25))
}
