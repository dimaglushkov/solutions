package main

import (
	"fmt"
	"github.com/dimaglushkov/solutions/adventofcode/util"
	"sort"
)

func checkReport(report []int) bool {
	if !sort.SliceIsSorted(report, func(i, j int) bool { return report[i] < report[j] }) &&
		!sort.SliceIsSorted(report, func(i, j int) bool { return report[i] > report[j] }) {
		return false
	}

	for i := 0; i < len(report)-1; i++ {
		if d := util.Abs(report[i] - report[i+1]); d == 0 || d > 3 {
			return false
		}
	}

	return true
}

func part1(reports [][]int) {
	var res int

	for _, report := range reports {
		if checkReport(report) {
			res++
		}

	}

	fmt.Println(res)
}

func part2(reports [][]int) {
	var res int

	for _, report := range reports {
		if checkReport(report) {
			res++
			continue
		}

		for i := range report {
			if checkReport(util.MergeSlices(report[:i], report[i+1:])) {
				res++
				break
			}
		}

	}

	fmt.Println(res)
}

func main() {
	var reports [][]int

	for _, l := range util.ReadInput(false) {
		reports = append(reports, util.ParseToIntArray(l, " "))
	}

	part1(reports)
	part2(reports)
}
