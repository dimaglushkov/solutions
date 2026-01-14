package main

import (
	"fmt"

	"github.com/dimaglushkov/solutions/adventofcode/util"
)

func part1(data []string) {
	sum := 50
	var ans int
	for _, d := range data {
		add := d[0] == 'R'
		val := util.Atoi(d[1:]) % 100

		if !add {
			val *= -1
		}

		sum += val
		if sum < 0 {
			sum += 100
		}
		if sum > 99 {
			sum -= 100
		}

		if sum == 0 {
			ans++
		}
	}

	fmt.Println(ans)
}

func part2(data []string) {
	sum := 50
	var ans int
	for _, d := range data {
		add := d[0] == 'R'
		val := util.Atoi(d[1:])

		if val > 100 {
			ans += val / 100
			val %= 100
		}
		if !add {
			val *= -1
		}

		if ns := sum + val; sum != 0 && (ns > 99 || ns < 0 || ns == 0) {
			ans++
		}

		sum += val
		if sum < 0 {
			sum += 100
		} else if sum > 99 {
			sum -= 100
		}
	}

	fmt.Println(ans)
}

func main() {
	data := util.ReadInput(false)

	//part1(data)
	part2(data)
}
