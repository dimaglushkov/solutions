package main

import (
	"fmt"
	"github.com/dimaglushkov/solutions/adventofcode/2024/util"
	"sort"
	"strconv"
	"strings"
)

const (
	//inputFile = "day01.example.input"
	inputFile = "day01.input"
)

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func part1(left, right []int) {
	var res int

	sort.Ints(left)
	sort.Ints(right)

	for i := 0; i < len(left); i++ {
		res += abs(left[i] - right[i])
	}

	fmt.Println(res)
}

func part2(left, right []int) {
	var res int

	cnt := make(map[int]int)
	for i := range right {
		cnt[right[i]]++
	}

	for i := range left {
		res += cnt[left[i]] * left[i]
	}

	fmt.Println(res)
}

func main() {
	lines := util.ReadInput(inputFile)

	var left, right []int
	var v1, v2 int
	var err error

	for i := range lines {
		s := strings.Split(lines[i], "   ")

		if v1, err = strconv.Atoi(s[0]); err != nil {
			panic(err)
		}

		if v2, err = strconv.Atoi(s[1]); err != nil {
			panic(err)
		}

		left = append(left, v1)
		right = append(right, v2)
	}

	part1(left, right)
	part2(left, right)

}
