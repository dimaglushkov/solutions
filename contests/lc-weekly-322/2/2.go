package main

func dividePlayers(skill []int) int64 {
	n := len(skill)
	s := 0

	m := make(map[int]int, n)
	for _, i := range skill {
		m[i]++
		s += i
	}
	ts := s / (n / 2)

	var res int64

	for _, x := range skill {
		if m[x] == 0 {
			continue
		}
		y := ts - x
		if m[y] != m[x] {
			return -1
		}
		c := m[x]
		if x == y {
			c /= 2
		}
		res += int64(c * (x * y))
		m[x], m[y] = 0, 0
	}

	return res
}

func main() {
	println(dividePlayers([]int{3, 2, 5, 1, 3, 4}))
	println(dividePlayers([]int{3, 4}))
	println(dividePlayers([]int{1, 1, 2, 3}))
}
