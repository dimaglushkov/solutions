package main

import (
	"fmt"
	"github.com/dimaglushkov/solutions/adventofcode/util"
	"regexp"
	"strings"
)

const (
	mulPattern       = `mul\(\d{1,3},\d{1,3}\)`
	mulDoDontPattern = `mul\(\d{1,3},\d{1,3}\)|do\(\)|don't\(\)`
)

func part1(input string) {
	var res int

	mulRe := regexp.MustCompile(mulPattern)
	matches := mulRe.FindAllStringSubmatch(input, -1)

	for _, m := range matches {
		x := strings.TrimPrefix(m[0], "mul(")
		x = strings.TrimSuffix(x, ")")

		values := strings.Split(x, ",")
		res += util.Atoi(values[0]) * util.Atoi(values[1])
	}

	fmt.Println(res)
}

func part2(input string) {
	var res int

	mulRe := regexp.MustCompile(mulDoDontPattern)
	matches := mulRe.FindAllStringSubmatch(input, -1)

	do := true

	for _, m := range matches {
		x := m[0]

		switch x {
		case "do()":
			do = true
		case "don't()":
			do = false
		default:
			if !do {
				break
			}
			x := strings.TrimPrefix(m[0], "mul(")
			x = strings.TrimSuffix(x, ")")

			values := strings.Split(x, ",")
			res += util.Atoi(values[0]) * util.Atoi(values[1])
		}
	}

	fmt.Println(res)
}

func main() {
	lines := util.ReadInput(false)

	part1(strings.Join(lines, "\n"))
	part2(strings.Join(lines, "\n"))
}
