package main

import (
	"fmt"
	"github.com/dimaglushkov/solutions/adventofcode/util"
	"strings"
)

func part1(rules map[int][]int, updates [][]int) [][]int {
	var res int
	var invalidUpdates [][]int

	for _, update := range updates {
		added := make(map[int]bool)
		set := util.SliceToSet(update)
		valid := true

		for _, x := range update {
			added[x] = true
			rulesX, ok := rules[x]
			if !ok {
				continue
			}

			for _, req := range rulesX {
				if _, ok := set[req]; !ok {
					continue
				}

				if _, ok := added[req]; !ok {
					valid = false
					break
				}
			}

			if !valid {
				break
			}
		}

		if valid {
			res += update[len(update)/2]
		} else {
			invalidUpdates = append(invalidUpdates, update)
		}

	}

	fmt.Println(res)
	return invalidUpdates
}

type queue []int

func (q *queue) push(x int) {
	*q = append(*q, x)
}
func (q *queue) pop() int {
	x := (*q)[0]
	*q = (*q)[1:]
	return x
}

func part2(rules map[int][]int, updates [][]int) {
	var res int

	for _, update := range updates {
		set := util.SliceToSet(update)
		q := queue{}
		fixedUpdate := make([]int, 0, len(update))

		filteredRules := make(map[int]map[int]bool, len(set))
		for x := range set {
			filteredRules[x] = make(map[int]bool, 0)
		}
		for k, v := range rules {
			if _, ok := filteredRules[k]; !ok {
				continue
			}
			for _, x := range v {
				if set[x] {
					filteredRules[k][x] = true
				}
			}
		}

		for k, v := range filteredRules {
			if len(v) == 0 {
				delete(filteredRules, k)
				q.push(k)
			}
		}

		for len(q) > 0 {
			x := q.pop()
			fixedUpdate = append(fixedUpdate, x)

			for r := range filteredRules {
				if _, ok := filteredRules[r][x]; ok {
					delete(filteredRules[r], x)
				}

				if len(filteredRules[r]) == 0 {
					q.push(r)
					delete(filteredRules, r)
				}
			}
		}

		res += fixedUpdate[len(fixedUpdate)/2]

	}

	fmt.Println(res)
}

func main() {
	lines := util.ReadInput(false)

	sep := 0
	for i := range lines {
		if lines[i] == "" {
			sep = i
			break
		}
	}

	rules := make(map[int][]int)
	for i := range lines[:sep] {
		vals := strings.Split(lines[i], "|")
		l, r := util.Atoi(vals[0]), util.Atoi(vals[1])

		if _, ok := rules[r]; !ok {
			rules[r] = make([]int, 0)
		}

		rules[r] = append(rules[r], l)
	}

	updates := make([][]int, 0)
	for _, l := range lines[sep+1:] {
		updateStr := strings.Split(l, ",")
		update := make([]int, len(updateStr))
		for i := range updateStr {
			update[i] = util.Atoi(updateStr[i])
		}
		updates = append(updates, update)
	}

	invalidUpdates := part1(rules, updates)

	part2(rules, invalidUpdates)
}
