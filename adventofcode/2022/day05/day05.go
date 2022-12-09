package main

import (
	"github.com/dimaglushkov/solutions/adventofcode/2022/util"
	"strconv"
	"strings"
)

func atoi(x string) int {
	y, _ := strconv.ParseInt(x, 10, 32)
	return int(y)
}

type stack []byte

func (s *stack) push(x byte) {
	*s = append(*s, x)
}

func (s *stack) pop() byte {
	n := len(*s)
	x := (*s)[n-1]
	*s = (*s)[:n-1]
	return x
}

func DayFive1() {
	text := util.ReadInput("day05_1.input")

	l := 0
	for text[l] != "" {
		l++
	}
	numbers := strings.Split(text[l-1], " ")
	n := atoi(numbers[len(numbers)-1])

	stacks := make([]stack, n)
	for i := l - 2; i >= 0; i-- {
		for j := 0; j < n && j*4+1 < len(text[i]); j++ {
			val := text[i][j*4+1]
			if val == ' ' {
				continue
			}
			stacks[j].push(val)
		}

	}

	for i := l + 1; i < len(text); i++ {
		vals := strings.Split(text[i], " ")
		cnt, from, to := atoi(vals[1]), atoi(vals[3])-1, atoi(vals[5])-1
		for j := 0; j < cnt; j++ {
			stacks[to].push(stacks[from].pop())
		}
	}

	res := ""
	for i := range stacks {
		res += string(stacks[i].pop())
	}

	println(res)
}

func DayFive2() {
	text := util.ReadInput("day05_1.input")

	l := 0
	for text[l] != "" {
		l++
	}
	numbers := strings.Split(strings.TrimRight(text[l-1], " "), " ")
	n := atoi(numbers[len(numbers)-1])

	stacks := make([][]byte, n)
	for i := range stacks {
		stacks[i] = make([]byte, 0)
	}

	for i := l - 2; i >= 0; i-- {
		for j := 0; j < n && j*4+1 < len(text[i]); j++ {
			val := text[i][j*4+1]
			if val == ' ' {
				continue
			}
			stacks[j] = append(stacks[j], val)
		}

	}

	for i := l + 1; i < len(text); i++ {
		vals := strings.Split(text[i], " ")
		cnt, from, to := atoi(vals[1]), atoi(vals[3])-1, atoi(vals[5])-1
		stacks[to] = append(stacks[to], stacks[from][len(stacks[from])-cnt:]...)
		stacks[from] = stacks[from][:len(stacks[from])-cnt]
		//for j := 0; j < cnt; j++ {
		//	stacks[to].push(stacks[from].pop())
		//}
	}

	res := ""
	for i := range stacks {
		res += string(stacks[i][len(stacks[i])-1])
	}

	println(res)
}

func main() {
	DayFive1()
	DayFive2()
}
