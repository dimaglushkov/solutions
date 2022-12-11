package main

import "github.com/dimaglushkov/solutions/adventofcode/2022/util"

func isSignal(x string) bool {
	n := len(x)
	for i := 0; i < n-1; i++ {
		for j := i + 1; j < n; j++ {
			if x[i] == x[j] {
				return false
			}
		}
	}
	return true
}

func DaySix1() {
	text := util.ReadInput("day06.input")[0]
	for i := 4; i < len(text); i++ {
		if isSignal(text[i-4 : i]) {
			println(i)
			return
		}
	}
}

func DaySix2() {
	text := util.ReadInput("day06.input")[0]
	for i := 14; i < len(text); i++ {
		if isSignal(text[i-14 : i]) {
			println(i)
			return
		}
	}
}

func main() {
	DaySix1()
	DaySix2()
}
