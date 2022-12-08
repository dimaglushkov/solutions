package day02

import (
	"github.com/dimaglushkov/solutions/adventofcode/2022/util"
	"strings"
)

func translate(r byte) byte {
	if r >= 'X' && r <= 'Z' {
		return r - 23
	}
	return r
}

// A - Rock
// B - Paper
// C - Scissors

func DayTwo1() {
	beats := map[byte]byte{
		'C': 'B',
		'B': 'A',
		'A': 'C',
	}

	res := 0
	for _, l := range util.ReadInput("day02_1.input") {
		rs := strings.Split(l, " ")
		a, b := rs[0][0], translate(rs[1][0])

		res += int(b - 'A' + 1)
		if a == b {
			res += 3
		} else if beats[b] == a {
			res += 6
		}
	}
	println(res)
}

func DayTwo2() {
	beats := map[byte]byte{
		'C': 'B',
		'B': 'A',
		'A': 'C',
	}
	loses := map[byte]byte{
		'B': 'C',
		'A': 'B',
		'C': 'A',
	}

	res := 0
	for _, l := range util.ReadInput("day02_1.input") {
		rs := strings.Split(l, " ")
		a, b := rs[0][0], rs[1][0]

		switch b {
		case 'X':
			res += int(beats[a] - 'A' + 1)

		case 'Y':
			res += 3
			res += int(a - 'A' + 1)

		case 'Z':
			res += 6
			res += int(loses[a] - 'A' + 1)
		}
	}

	println(res)
}
