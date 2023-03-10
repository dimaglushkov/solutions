package main

import (
	"fmt"
	. "github.com/dimaglushkov/solutions/ADS/list"
	"log"
	"math/rand"
	"os"
	"time"
)

// source: https://leetcode.com/problems/linked-list-random-node/

type Solution struct {
	head *ListNode
	r    *rand.Rand
}

func Constructor(head *ListNode) Solution {
	return Solution{
		head: head,
		r:    rand.New(rand.NewSource(time.Now().UnixNano())),
	}
}

func (this *Solution) GetRandom() int {
	ans, n := 0, 0
	for p := this.head; p != nil; p = p.Next {
		n++
		if this.r.Intn(n) == 0 {
			ans = p.Val
		}
	}
	return ans
}

/**
 * Your Solution object will be instantiated and called as such:
 * obj := Constructor(head);
 * param_1 := obj.GetRandom();
 */

func main() {
	size := 1000
	test := 100_000
	values := make([]int, size)
	cnt := make([]int, size)
	for i := 0; i < size; i++ {
		values[i] = i
	}

	sol := Constructor(NewList(values))
	for i := 0; i < test; i++ {
		cnt[sol.GetRandom()]++
	}
	f, err := os.OpenFile("/tmp/stats.csv", os.O_RDWR, 0755)
	if err != nil {
		log.Fatal(err)
	}
	for i := range cnt {
		if _, err := f.WriteString(fmt.Sprintf("%d,%d\n", i, cnt[i])); err != nil {
			log.Fatal(err)
		}
	}
	f.Close()
}
