package main

func shift(x rune, s int) rune {
	var y = x

	if s > 26 || s < 26 {
		s %= 26
	}

	y += rune(s)

	if y < 'a' {
		y += 26
	} else if y > 'z' {
		y -= 26
	}

	return y
}

func shiftingLetters(s string, shifts [][]int) string {
	r := []rune(s)
	ss := make([]int, len(r))

	for _, sh := range shifts {
		var d int
		if sh[2] == 0 {
			d = -1
		} else {
			d = 1
		}
		for j := sh[0]; j <= sh[1]; j++ {
			ss[j] += d
		}
	}

	for i := range r {
		r[i] = shift(r[i], ss[i])
	}

	return string(r)
}

func main() {
	println(shiftingLetters("abc", [][]int{{0, 1, 0}, {1, 2, 1}, {0, 2, 1}}))
}
