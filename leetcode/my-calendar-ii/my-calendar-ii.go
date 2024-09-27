package main

import "sort"

// source: https://leetcode.com/problems/my-calendar-ii/
type pair struct {
	time  int
	delta int
}

type MyCalendarTwo struct {
	entries []pair
}

func Constructor() MyCalendarTwo {
	return MyCalendarTwo{
		entries: make([]pair, 0),
	}
}

func (mct *MyCalendarTwo) Book(start int, end int) bool {
	var cnt int
	var ans = true

	mct.entries = append(mct.entries, pair{time: start, delta: 1})
	mct.entries = append(mct.entries, pair{time: end, delta: -1})
	sort.Slice(mct.entries, func(i, j int) bool {
		if mct.entries[i].time == mct.entries[j].time {
			return mct.entries[i].delta < mct.entries[j].delta
		}
		return mct.entries[i].time < mct.entries[j].time
	})

	for i := 0; i < len(mct.entries); i++ {
		cnt += mct.entries[i].delta

		if cnt > 2 {
			ans = false
		}
	}

	if !ans {
		var startChanged, endChanged bool
		for i := range mct.entries {
			if !startChanged && mct.entries[i].time == start && mct.entries[i].delta == 1 {
				mct.entries[i].delta = 0
				startChanged = true
			}
			if !endChanged && mct.entries[i].time == end && mct.entries[i].delta == -1 {
				mct.entries[i].delta = 0
				endChanged = true
			}
		}
	}

	return ans
}

func main() {
	mct := Constructor()

	println(mct.Book(10, 20))
	println(mct.Book(50, 60))
	println(mct.Book(10, 40))
	println(mct.Book(5, 15))
	println(mct.Book(5, 10))
	println(mct.Book(25, 55))
	println()

	mct2 := Constructor()

	println(mct2.Book(26, 35))
	println(mct2.Book(26, 32))
	println(mct2.Book(25, 32))
	println(mct2.Book(18, 26))
	println(mct2.Book(40, 45))
	println(mct2.Book(19, 26))
	println(mct2.Book(48, 50))
	println(mct2.Book(1, 6))
	println(mct2.Book(46, 50))
	println(mct2.Book(11, 18))

}

/**
 * Your MyCalendarTwo object will be instantiated and called as such:
 * obj := Constructor();
 * param_1 := obj.Book(start,end);
 */
