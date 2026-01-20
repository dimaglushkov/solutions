package main

import (
	"fmt"
	"strings"

	"github.com/dimaglushkov/solutions/adventofcode/util"
)

const (
	add = iota
	mul
)

func process(input []string) ([][]int, []int) {
	var ops []int

	for _, op := range strings.Split(input[len(input)-1], " ") {
		if op == "+" {
			ops = append(ops, add)
		}
		if op == "*" {
			ops = append(ops, mul)
		}
	}

	values := make([][]int, len(input)-1)
	for i := range input[:len(input)-1] {
		values[i] = make([]int, 0, len(ops))
		for _, val := range strings.Split(input[i], " ") {
			if val == "" {
				continue
			}

			values[i] = append(values[i], util.Atoi(val))
		}
	}

	return values, ops
}

func part1(input []string) {
	var res int

	vals, ops := process(input)

	totals := vals[0]
	vals = vals[1:]

	for i := range vals {
		for j := range vals[i] {
			switch ops[j] {
			case add:
				totals[j] += vals[i][j]
			case mul:
				totals[j] *= vals[i][j]
			}
		}
	}

	for i := range totals {
		res += totals[i]
	}

	fmt.Println(res)
}

func part2(input []string) {
	var res int

	var starts []int
	var ops []int

	for i, v := range input[len(input)-1] {
		if v != ' ' {
			starts = append(starts, i)
		}

		switch v {
		case '+':
			ops = append(ops, add)
		case '*':
			ops = append(ops, mul)
		}
	}

	input = input[:len(input)-1]

	maxLen := len(input[0])
	for i := range input {
		maxLen = max(maxLen, len(input[i]))
	}

	for i := range input {
		for len(input[i]) < maxLen {
			input[i] += " "
		}
	}

	for i := range starts {
		r := maxLen
		if i < len(starts)-1 {
			r = starts[i+1] - 1
		}

		var total int
		if ops[i] == mul {
			total = 1
		}

		for l := starts[i]; l < r; l++ {
			var num int
			for j := 0; j < len(input); j++ {
				if input[j][l] == ' ' {
					continue
				}
				num *= 10
				num += util.Atoi(string(input[j][l]))
			}

			switch ops[i] {
			case add:
				total += num
			case mul:
				total *= num
			}
		}

		res += total
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
