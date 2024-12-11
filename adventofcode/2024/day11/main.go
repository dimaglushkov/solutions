package main

import (
	"fmt"
	"github.com/dimaglushkov/solutions/adventofcode/util"
	"math/big"
	"strconv"
	"strings"
)

type List struct {
	val  *big.Int
	next *List
}

func part1(head *List) {
	x0 := new(big.Int)
	x2024 := big.NewInt(2024)

	for range 25 {
		for c := head; c != nil; c = c.next {
			switch {
			case x0.Cmp(c.val) == 0:
				c.val = big.NewInt(1)
			case len(c.val.Text(10))%2 == 0:
				str := c.val.Text(10)
				c.val, _ = new(big.Int).SetString(str[:len(str)/2], 10)

				tmp, _ := new(big.Int).SetString(strings.TrimPrefix(str[len(str)/2:], "0"), 10)
				if tmp == nil {
					tmp = new(big.Int)
				}

				next := &List{
					val:  tmp,
					next: c.next,
				}
				c.next = next
				c = c.next
			default:
				c.val = c.val.Mul(c.val, x2024)
			}
		}
	}

	var res int

	for head != nil {
		res++
		head = head.next
	}

	fmt.Println(res)
}

func part2(nums []int) {
	var res int

	cnt := make(map[int]int)
	for _, x := range nums {
		cnt[x]++
	}

	for range 75 {
		newCnt := make(map[int]int)
		for x, c := range cnt {
			s := strconv.Itoa(x)
			switch {
			case x == 0:
				newCnt[1] += c
			case len(s)%2 == 0:
				newCnt[util.Atoi(s[:len(s)/2])] += c
				newCnt[util.Atoi(s[len(s)/2:])] += c
			default:
				newCnt[x*2024] += c
			}
		}
		cnt = newCnt
	}

	for _, c := range cnt {
		res += c
	}

	fmt.Println(res)
}

func main() {
	lines := util.ReadInput(false)

	head := new(List)
	cur := head

	for _, i := range strings.Split(lines[0], " ") {
		x, _ := new(big.Int).SetString(i, 10)
		cur.val = x
		cur.next = new(List)
		cur = cur.next
	}

	for cur = head; cur.next.next != nil; cur = cur.next {
	}

	cur.next = nil

	part1(head)

	var nums []int
	for _, i := range strings.Split(lines[0], " ") {
		nums = append(nums, util.Atoi(i))
	}

	part2(nums)
}
