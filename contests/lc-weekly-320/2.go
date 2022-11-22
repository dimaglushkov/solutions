package main

import (
	"sort"

	. "github.com/dimaglushkov/solutions/ADS/tree"
)

// TLE ???
// maybe should've used cycles instead of recursion
// anyway spent too much time on this simple problem :/

//func closestNodes(root *TreeNode, queries []int) [][]int {
//	res := make([][]int, 0, len(queries))
//	var maxVal, minVal int
//	var find func(n *TreeNode, val int)
//	find = func(n *TreeNode, val int) {
//		if n == nil {
//			return
//		}
//		if n.Val >= val && n.Val < maxVal {
//			maxVal = n.Val
//		}
//		if n.Val <= val && n.Val > minVal {
//			minVal = n.Val
//		}
//		if val < n.Val {
//			find(n.Left, val)
//		} else {
//			find(n.Right, val)
//		}
//	}
//
//	for _, q := range queries {
//		var min, max = -1, -1
//		minVal = math.MinInt
//		maxVal = math.MaxInt
//		find(root, q)
//		if minVal != math.MinInt {
//			min = minVal
//		}
//		if maxVal != math.MaxInt {
//			max = maxVal
//		}
//
//		res = append(res, []int{min, max})
//	}
//
//	return res
//}

func closestNodes(root *TreeNode, queries []int) [][]int {
	d := make([]int, 0)
	var getData func(n *TreeNode)
	getData = func(n *TreeNode) {
		if n == nil {
			return
		}
		d = append(d, n.Val)
		getData(n.Left)
		getData(n.Right)
	}
	getData(root)
	sort.Ints(d)

	res := make([][]int, 0, len(queries))
	for _, q := range queries {
		var min, max int = -1, -1
		x := sort.SearchInts(d, q)
		if x >= len(d) {
			min = d[x-1]
		} else if d[x] == q {
			min, max = d[x], d[x]
		} else if x == 0 {
			max = d[x]
		} else {
			min, max = d[x-1], d[x]
		}

		res = append(res, []int{min, max})
	}
	return res
}

func main() {
	println(closestNodes(NewTreeNode([]int{6, 2, 13, 1, 4, 9, 15, -1, -1, -1, -1, -1, -1, 14}), []int{0, 2, 5, 16}))
	println(closestNodes(NewTreeNode([]int{4, -1, 9}), []int{3}))
	println(closestNodes(NewTreeNode([]int{16, 8, 18, 1, 12, -1, 20, -1, 2, 9, -1, -1, -1, -1, 7}), []int{6}))
	println(closestNodes(NewTreeNode([]int{16, 8, 18, 1, 12, -1, 20, -1, 2, 9, -1, -1, -1, -1, 7}), []int{6}))

}
