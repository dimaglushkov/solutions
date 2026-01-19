package main

import (
	"fmt"
	"sort"
	"strings"

	"github.com/dimaglushkov/solutions/adventofcode/util"
)

func process(input []string) ([][]int, []int) {
	intervals := make([][]int, 0)
	ids := make([]int, 0)

	i := 0
	for input[i] != "" {
		numbers := strings.Split(input[i], "-")
		intervals = append(intervals, []int{util.Atoi(numbers[0]), util.Atoi(numbers[1])})
		i++
	}

	i++

	for i < len(input) {
		ids = append(ids, util.Atoi(input[i]))
		i++
	}

	return intervals, ids
}

func mergeIntervals(intervals [][]int) [][]int {
	entries := make([]int, 0, len(intervals)*2)
	for i := range intervals {
		entries = append(entries, intervals[i][0], intervals[i][1]*-1)
	}

	sort.Slice(entries, func(i, j int) bool {
		return util.Abs(entries[i]) < util.Abs(entries[j])
	})

	var c, l int
	mergedIntervals := make([][]int, 0)
	for i := range entries {
		if c == 0 {
			l = entries[i]
		}

		if entries[i] > 0 {
			c++
		} else {
			c--
		}

		if c == 0 {
			mergedIntervals = append(mergedIntervals, []int{util.Abs(l), util.Abs(entries[i])})
		}
	}

	return mergedIntervals
}

func part1(input []string) {
	var res int

	intervals, ids := process(input)
	intervals = mergeIntervals(intervals)

	for _, id := range ids {
		x := sort.Search(len(intervals), func(i int) bool {
			return intervals[i][1] >= id
		})

		if x >= 0 && x < len(intervals) && intervals[x][0] <= id && intervals[x][1] >= id {
			res++
		}
	}

	fmt.Println(res)
}

func part2(input []string) {
	var res int

	intervals, _ := process(input)
	intervals = mergeIntervals(intervals)

	counted := make(map[int]int) // workaround to avoid double counting [a - X] and [X - b] as well as [X - X]

	for _, interval := range intervals {
		if interval[0] == interval[1] {
			res++
		} else {
			res += interval[1] - interval[0] + 1
		}

		if counted[interval[0]] > 0 {
			res--
		}

		counted[interval[0]]++
		counted[interval[1]]++

	}

	fmt.Println(res)
}

func main() {
	exampleLines := util.ReadInput(true)

	fmt.Println("Example:")
	part1(exampleLines)
	part2(exampleLines)

	lines := util.ReadInput(false)
	fmt.Println("=========")
	part1(lines)
	part2(lines)
}
