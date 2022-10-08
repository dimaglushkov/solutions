package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/my-calendar-iii/

type MyCalendarThree struct {
	cnt    int
	starts []int
	ends   []int
}

func Constructor() MyCalendarThree {
	return MyCalendarThree{
		cnt:    0,
		starts: make([]int, 0),
		ends:   make([]int, 0),
	}
}

func (this *MyCalendarThree) Book(start int, end int) int {
	n := len(this.starts)
	si := sort.Search(n, func(i int) bool {
		return this.starts[i] >= start
	})
	this.starts = append(this.starts, 0)
	copy(this.starts[si+1:], this.starts[si:])
	this.starts[si] = start

	ei := sort.Search(n, func(i int) bool {
		return this.ends[i] >= end
	})
	this.ends = append(this.ends, 0)
	copy(this.ends[ei+1:], this.ends[ei:])
	this.ends[ei] = end

	n++
	var i, j, cnt int
	for i = 0; i < n && this.starts[i] <= start; i++ {
		cnt++
	}
	for j = 0; j < n && this.ends[j] <= start; j++ {
		cnt--
	}
	if cnt > this.cnt {
		this.cnt = cnt
	}
	for (i < n && this.starts[i] <= end) && (j < n && this.ends[j] <= end) {
		if this.ends[j] <= this.starts[i] {
			cnt--
			j++
		} else {
			cnt++
			i++
			if cnt > this.cnt {
				this.cnt = cnt
			}
		}
	}
	for i < n && this.starts[i] <= end {
		cnt++
		i++
	}
	if cnt > this.cnt {
		this.cnt = cnt
	}

	return this.cnt
}

func main() {
	x := Constructor()

	//1,1,2,2,3,3,3,3,4,4
	//1,1,2,2,2,4,4,4,4,4
	//print(x.Book(24, 40))
	//print(x.Book(43, 50))
	//print(x.Book(27, 43))
	//print(x.Book(5, 21))
	//print(x.Book(30, 40))
	//print(x.Book(14, 29)) // 1, 3, 4
	//print(x.Book(3, 19))
	//print(x.Book(3, 14))
	//print(x.Book(25, 39))
	//print(x.Book(6, 19))

	//1,1,1,1,1,2,2,2,3,3,3,4,5,5,5,5,5,5,6,6,6,6,6,6,7,7,7,7,7,7
	fmt.Printf("%d,", x.Book(47, 50))
	fmt.Printf("%d,", x.Book(1, 10))
	fmt.Printf("%d,", x.Book(27, 36))
	fmt.Printf("%d,", x.Book(40, 47))
	fmt.Printf("%d,", x.Book(20, 27))
	fmt.Printf("%d,", x.Book(15, 23))
	fmt.Printf("%d,", x.Book(10, 18))
	fmt.Printf("%d,", x.Book(27, 36))
	fmt.Printf("%d,", x.Book(17, 25))
	fmt.Printf("%d,", x.Book(8, 17))
	fmt.Printf("%d,", x.Book(24, 33))
	fmt.Printf("%d,", x.Book(23, 28))
	fmt.Printf("%d,", x.Book(21, 27))
	fmt.Printf("%d,", x.Book(47, 50))
	fmt.Printf("%d,", x.Book(14, 21))
	fmt.Printf("%d,", x.Book(26, 32))
	fmt.Printf("%d,", x.Book(16, 21))
	fmt.Printf("%d,", x.Book(2, 7))
	fmt.Printf("%d,", x.Book(24, 33))
	fmt.Printf("%d,", x.Book(6, 13))
	fmt.Printf("%d,", x.Book(44, 50))
	fmt.Printf("%d,", x.Book(33, 39))
	fmt.Printf("%d,", x.Book(30, 36))
	fmt.Printf("%d,", x.Book(6, 15))
	fmt.Printf("%d,", x.Book(21, 27))
	fmt.Printf("%d,", x.Book(49, 50))
	fmt.Printf("%d,", x.Book(38, 45))
	fmt.Printf("%d,", x.Book(4, 12))
	fmt.Printf("%d,", x.Book(46, 50))
	fmt.Printf("%d,", x.Book(13, 21))

	//println(x.Book(10, 20))
	//println(x.Book(50, 60))
	//println(x.Book(10, 40))
	//println(x.Book(5, 15))
	//println(x.Book(5, 10))
	//println(x.Book(25, 55))

	//println(x.Book(26, 35))
	//println(x.Book(26, 32))
	//println(x.Book(25, 32))
	//println(x.Book(18, 26))
	//println(x.Book(40, 45))
	//println(x.Book(19, 26))
	//println(x.Book(48, 50))
	//println(x.Book(1, 6))
	//println(x.Book(46, 50))
	//println(x.Book(11, 18))
}
