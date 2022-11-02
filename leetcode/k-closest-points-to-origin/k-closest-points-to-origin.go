package main

import (
	"fmt"
	"math"
	"sort"
)

// source: https://leetcode.com/problems/k-closest-points-to-origin/

func kClosest(points [][]int, k int) [][]int {
	dist := func(x, y int) float64 {
		xf, yf := float64(x), float64(y)
		return math.Sqrt(xf*xf + yf*yf)
	}

	type p struct {
		xy []int
		d  float64
	}
	ps := make([]p, 0)
	for _, i := range points {
		ps = append(ps, p{xy: i, d: dist(i[0], i[1])})
	}

	sort.Slice(ps, func(i, j int) bool {
		return ps[i].d < ps[j].d
	})

	res := make([][]int, k)
	for i := range res {
		res[i] = ps[i].xy
	}
	return res
}

func main() {
	var points0 = [][]int{{-5, 4}, {-6, -5}, {4, 6}}
	var k0 int = 2
	fmt.Println("Expected: [[-5,4],[4,6]]	Output: ", kClosest(points0, k0))

	// Example 1
	var points1 = [][]int{{1, 3}, {-2, 2}}
	var k1 int = 1
	fmt.Println("Expected: [[-2,2]]	Output: ", kClosest(points1, k1))

	// Example 2
	var points2 = [][]int{{3, 3}, {5, -1}, {-2, 4}}
	var k2 int = 2
	fmt.Println("Expected: [[3,3],[-2,4]]	Output: ", kClosest(points2, k2))

}
