package main

import (
	"fmt"

	"github.com/dimaglushkov/solutions/adventofcode/util"
)

func part1(input [][]int) {
	var res int

	for _, bank := range input {
		var maxVal, maxId int

		for i := range bank {
			if bank[i] > maxVal {
				maxVal = bank[i]
				maxId = i
			}
		}

		var j int

		if maxId != len(bank)-1 {
			sMaxVal := 0
			for _, v := range bank[maxId+1:] {
				sMaxVal = max(sMaxVal, v)
			}

			j = maxVal*10 + sMaxVal
		} else {
			sMaxVal := 0
			for _, v := range bank[:maxId] {
				sMaxVal = max(sMaxVal, v)
			}

			j = sMaxVal*10 + maxVal
		}

		res += j
	}

	fmt.Println(res)
}

func part2(input [][]int) {
	var res int

	for _, bank := range input {
		x := 0
		l, r := 0, len(bank)-11

		for range 12 {
			var mv, mi int
			for i := l; i < r; i++ {
				if bank[i] > mv {
					mv = bank[i]
					mi = i
				}
			}

			x = x*10 + mv
			r++
			l = mi + 1
		}

		res += x
	}

	fmt.Println(res)
}

//func part2(input [][]int) {
//	var res int
//
//	for _, bank := range input {
//		used := make([]bool, len(bank))
//
//		leftmost := -1
//		for i := 9; i >= 0 && leftmost == -1; i-- {
//			for j := 0; j < len(bank); j++ {
//				if bank[j] == i && len(bank)-j > 12 {
//					leftmost = j
//					break
//				}
//			}
//		}
//
//		for dl := 12; dl > 0; dl-- {
//			var mv, mi int
//			for i := len(bank) - 1; i >= leftmost; i-- {
//				if !used[i] && bank[i] > mv {
//					mv = bank[i]
//					mi = i
//				}
//			}
//
//			used[mi] = true
//		}
//
//		x := 0
//		for i := range used {
//			if used[i] {
//				x = x*10 + bank[i]
//			}
//		}
//
//		res += x
//	}
//
//	fmt.Println(res)
//}

func main() {
	exampleLines := util.ReadInput(true)

	exampleData := make([][]int, len(exampleLines))
	for i, l := range exampleLines {
		exampleData[i] = util.ParseToIntArray(l, "")
	}

	fmt.Println("Example:")
	part1(exampleData)
	part2(exampleData)

	lines := util.ReadInput(false)

	data := make([][]int, len(lines))
	for i, l := range lines {
		data[i] = util.ParseToIntArray(l, "")
	}

	fmt.Println("=========")
	part1(data)
	part2(data)
}
