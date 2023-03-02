package convex_hull

import (
	"math"
	"sort"
)

func direction(a, b, c []int) int {
	return (b[0]-a[0])*(c[1]-a[1]) - (b[1]-a[1])*(c[0]-a[0])
}

func distance(a, b []int) float64 {
	d1, d2 := float64(a[0]-b[0]), float64(a[1]-b[1])
	return math.Sqrt(d1*d1 + d2*d2)
}

// Jarvis This implementation doesn't handle collinear cases properly
func Jarvis(points [][]int) [][]int {
	n := len(points)
	if n <= 3 {
		return points
	}

	start := 0
	for i := 1; i < n; i++ {
		if points[i][1] < points[start][1] || points[i][1] == points[start][1] && points[i][0] < points[start][0] {
			start = i
		}
	}

	res := make([][]int, 0)
	cur := start
	v := make([]bool, n)
	v[start] = true
	for {
		c := -1
		for i := 0; i < n; i++ {
			if i == cur {
				continue
			}
			if c == -1 {
				c = i
				continue
			}
			d := direction(points[cur], points[c], points[i])
			if d == 0 && distance(points[cur], points[i]) < distance(points[cur], points[c]) {
				c = i
			} else if d < 0 {
				c = i
			}
		}

		res = append(res, points[c])
		if c == start {
			break
		}
		v[cur] = true
		cur = c
	}

	return res
}

//func GrahamScan(points [][]int) [][]int {
//	if len(points) <= 3 {
//		return points
//	}
//
//	n := len(points)
//
//	start := 0
//	for i := 1; i < n; i++ {
//		if points[i][1] < points[start][1] || points[i][1] == points[start][1] && points[i][0] < points[start][0] {
//			start = i
//		}
//	}
//
//	sort.Slice(points, func(i, j int) bool {
//		if d := direction(points[start], points[i], points[j]); d != 0 {
//			return d < 0
//		}
//		return distance(points[start], points[i]) < distance(points[start], points[j])
//	})
//
//	return nil
//}

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
