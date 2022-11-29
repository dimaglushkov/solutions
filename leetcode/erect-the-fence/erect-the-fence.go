package main

import (
	"fmt"
	"sort"
)

// source: https://leetcode.com/problems/erect-the-fence/

//func direction(a, b, c []int) int {
//	return (b[0]-a[0])*(c[1]-a[1]) - (b[1]-a[1])*(c[0]-a[0])
//}
//
//func dist(a, b []int) float64 {
//	d1, d2 := float64(a[0]-b[0]), float64(a[1]-b[1])
//	return math.Sqrt(d1*d1 + d2*d2)
//}
//
//// Implementing Jarvis march algorithm
//func outerTrees(trees [][]int) [][]int {
//	n := len(trees)
//	if n <= 3 {
//		return trees
//	}
//
//	start := 0
//	for i := 1; i < n; i++ {
//		if trees[i][1] < trees[start][1] || trees[i][1] == trees[start][1] && trees[i][0] < trees[start][0] {
//			start = i
//		}
//	}
//
//	res := make([][]int, 0)
//	cur := start
//	prevDir := true
//	v := make([]bool, n)
//	v[start] = true
//	for {
//		c := -1
//		for i := 0; i < n; i++ {
//			if i == cur {
//				continue
//			}
//			if c == -1 {
//				c = i
//				continue
//			}
//			d := direction(trees[cur], trees[c], trees[i])
//			if d == 0 && dist(trees[cur], trees[i]) < dist(trees[cur], trees[c]) {
//				if col :=
//				c = i
//				prevDir = dir
//			} else if d < 0 {
//				c = i
//			}
//		}
//
//		res = append(res, trees[c])
//		if c == start {
//			break
//		}
//		v[cur] = true
//		cur = c
//	}
//
//	return res
//}

func compare(a, b, c []int) int {
	return (c[1]-b[1])*(b[0]-a[0]) - (b[1]-a[1])*(c[0]-b[0])
}

func outerTrees(trees [][]int) [][]int {
	if len(trees) <= 3 {
		return trees
	}

	sort.Slice(trees, func(i, j int) bool {
		if trees[i][0]-trees[j][0] == 0 {
			return trees[i][1] < trees[j][1]
		}
		return trees[i][0] < trees[j][0]
	})

	u, l := make([][]int, 0), make([][]int, 0)
	for i := range trees {
		for len(l) >= 2 && compare(l[len(l)-2], l[len(l)-1], trees[i]) > 0 {
			l = l[:len(l)-1]
		}
		for len(u) >= 2 && compare(u[len(u)-2], u[len(u)-1], trees[i]) < 0 {
			u = u[:len(u)-1]
		}
		l = append(l, trees[i])
		u = append(u, trees[i])
	}

	res := make([][]int, 0)
	m := make(map[int]bool)
	for _, s := range append(u, l...) {
		k := s[0]*1000 + s[1]
		if m[k] {
			continue
		}
		res = append(res, s)
		m[k] = true
	}

	return res
}

func main() {
	testCases := []struct {
		trees [][]int
		want  [][]int
	}{
		{
			trees: [][]int{{1, 1}, {2, 2}, {2, 0}, {2, 4}, {3, 3}, {4, 2}},
			want:  [][]int{{1, 1}, {2, 0}, {2, 4}, {3, 3}, {4, 2}},
		},
		{
			trees: [][]int{{1, 2}, {2, 2}, {4, 2}, {5, 2}, {6, 2}, {7, 2}},
			want:  [][]int{{4, 2}, {6, 2}, {2, 2}, {5, 2}, {1, 2}, {7, 2}},
		},

		{
			trees: [][]int{{1, 2}, {2, 2}, {4, 2}},
			want:  [][]int{{1, 2}, {2, 2}, {4, 2}},
		},
	}

	successes := 0
	for _, tc := range testCases {
		x := outerTrees(tc.trees)
		sort.Slice(x, func(i, j int) bool {
			if x[i][0] == x[j][0] {
				return x[i][1] < x[j][1]
			}
			return x[i][0] < x[j][0]
		})
		sort.Slice(tc.want, func(i, j int) bool {
			if tc.want[i][0] == tc.want[j][0] {
				return tc.want[i][1] < tc.want[j][1]
			}
			return tc.want[i][0] < tc.want[j][0]
		})

		status := "ERROR"
		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
			status = "OK"
			successes++
		}
		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
	}
	if l := len(testCases); successes == len(testCases) {
		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
	} else {
		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
	}

}
