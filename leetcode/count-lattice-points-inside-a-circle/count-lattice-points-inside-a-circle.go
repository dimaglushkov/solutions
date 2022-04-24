package main

func countLatticePoints(circles [][]int) (res int) {
	points := make(map[[2]int]bool)

	for _, circle := range circles {
		x, y, r := circle[0], circle[1], circle[2]
		for i := x - r; i <= x+r; i++ {
			for j := y - r; j <= y+r; j++ {
				if (x-i)*(x-i)+(y-j)*(y-j) <= r*r {
					points[[2]int{i, j}] = true
				}
			}
		}
	}

	for _ = range points {
		res++
	}
	return res
}

func main() {
	println(countLatticePoints([][]int{{2, 2, 1}}))
	println(countLatticePoints([][]int{{2, 2, 2}, {3, 4, 1}}))
}
