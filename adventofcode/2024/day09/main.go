package main

import (
	"fmt"
	"github.com/dimaglushkov/solutions/adventofcode/util"
)

const space = -1

func decompress(s string) []int {
	res := make([]int, 0)
	id := 0
	for i := range s {
		x := util.Atoi(string(s[i]))
		if i%2 == 0 {
			for j := 0; j < x; j++ {
				res = append(res, id)
			}
			id++
		} else {
			for j := 0; j < x; j++ {
				res = append(res, space)
			}
		}
	}

	return res
}

func part1(mem []int) {
	var res int

	l, r := 0, len(mem)-1

	for l < r {
		for mem[l] != space {
			l++
		}
		for mem[r] == space {
			r--
		}

		mem[l], mem[r] = mem[r], mem[l]

		l++
		r--
	}

	for i := range mem {
		if mem[i] == space {
			continue
		}

		res += i * mem[i]
	}

	fmt.Println(res)
}

// this is probably the ugliest piece of code I've written in the year
func part2(mem []int) {
	var res int

	r := len(mem) - 1

	for r > 0 {
		for mem[r] == space {
			r--
		}
		flen := 0
		fid := mem[r]
		for r > 0 && mem[r] == fid {
			flen++
			r--
		}
		if r != 0 {
			r++
		}

		l, slen := 0, 0
		for l < len(mem) {
			for l < len(mem) && mem[l] != space {
				l++
			}

			slen = 0
			for l+slen < len(mem) && mem[l+slen] == space {
				slen++
			}

			if slen >= flen {
				break
			}

			l++
		}

		if l == len(mem) || flen > slen || r < l {
			r--
			continue
		}

		for i := 0; i < flen; i++ {
			mem[l+i] = fid
			mem[r+i] = space
		}
		r--

		//printMem(mem)
	}

	for i := range mem {
		if mem[i] == space {
			continue
		}

		res += i * mem[i]
	}
	fmt.Println(res)
}

func printMem(mem []int) {
	for i := range mem {
		if mem[i] == space {
			fmt.Print(".")
		} else {
			fmt.Printf("%d", mem[i])
		}
	}
	fmt.Println()
}

func main() {
	lines := util.ReadInput(false)

	mem := decompress(lines[0])
	mem2 := make([]int, len(mem))
	copy(mem2, mem)

	part1(mem)
	part2(mem2)
}
