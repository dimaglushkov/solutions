package day03

import "github.com/dimaglushkov/solutions/adventofcode/2022/util"

func prior(x byte) int {
	if x >= 'a' && x <= 'z' {
		return int(x-'a') + 1
	}
	if x >= 'A' && x <= 'Z' {
		return int(x-'A') + 27
	}
	return 0
}

func DayThree1() {
	res := 0
	for _, l := range util.ReadInput("day03_1.input") {
		n := len(l)
		m1, m2 := make(map[byte]bool), make(map[byte]bool)
		for i := 0; i < n/2; i++ {
			m1[l[i]] = true
		}
		for i := n / 2; i < n; i++ {
			m2[l[i]] = true
		}

		for k := range m1 {
			if m2[k] {
				res += prior(k)
				break
			}
		}

	}
	println(res)
}

func DayThree2() {
	const (
		groupSize = 3
		priorCnt  = 53
	)
	res := 0
	lines := util.ReadInput("day03_1.input")
	i, n := 0, len(lines)
	for i < n {
		var cnt [groupSize][priorCnt]int
		for j := 0; j < groupSize; j++ {
			for _, x := range lines[i] {
				cnt[j][prior(byte(x))]++
			}
			i++
		}

		for j := 0; j < priorCnt; j++ {
			inAllGroups := true
			for k := 0; k < groupSize; k++ {
				if cnt[k][j] == 0 {
					inAllGroups = false
					break
				}
			}
			if inAllGroups {
				res += j
			}
		}
	}

	println(res)
}
