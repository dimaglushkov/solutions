package convex_hull

import (
	"sort"
)

func direction(a, b, c []int) int {
	return (b[0]-a[0])*(c[1]-a[1]) - (b[1]-a[1])*(c[0]-a[0])
}

func MonotonicChain(points [][]int) [][]int {
	if len(points) <= 3 {
		return points
	}

	sort.Slice(points, func(i, j int) bool {
		if points[i][0]-points[j][0] == 0 {
			return points[i][1] < points[j][1]
		}
		return points[i][0] < points[j][0]
	})

	upper, lower := make([][]int, 0), make([][]int, 0)
	for i := range points {
		for len(lower) >= 2 && direction(lower[len(lower)-2], lower[len(lower)-1], points[i]) > 0 {
			lower = lower[:len(lower)-1]
		}
		for len(upper) >= 2 && direction(upper[len(upper)-2], upper[len(upper)-1], points[i]) < 0 {
			upper = upper[:len(upper)-1]
		}
		lower = append(lower, points[i])
		upper = append(upper, points[i])
	}

	res := make([][]int, 0)
	m := make(map[int]bool)
	for _, s := range append(upper, lower...) {
		k := s[0]*1000 + s[1]
		if m[k] {
			continue
		}
		res = append(res, s)
		m[k] = true
	}

	return res
}
