package main

// source: https://leetcode.com/problems/construct-quad-tree/

type Node struct {
	Val         bool
	IsLeaf      bool
	TopLeft     *Node
	TopRight    *Node
	BottomLeft  *Node
	BottomRight *Node
}

func construct(grid [][]int) *Node {
	n := len(grid)
	boolMap := [2]bool{false, true}

	var build func(i, j, l int) *Node
	build = func(i, j, l int) *Node {
		if l == 0 {
			return nil
		}

		if l == 1 {
			return &Node{Val: boolMap[grid[i][j]], IsLeaf: true}
		}

		tl, tr := build(i, j, l/2), build(i, j+l/2, l/2)
		bl, br := build(i+l/2, j, l/2), build(i+l/2, j+l/2, l/2)
		if tl.IsLeaf && tr.IsLeaf && bl.IsLeaf && br.IsLeaf &&
			tl.Val == tr.Val && tr.Val == bl.Val && bl.Val == br.Val {
			return &Node{Val: tl.Val, IsLeaf: true}
		}

		return &Node{
			IsLeaf:      false,
			TopLeft:     tl,
			TopRight:    tr,
			BottomLeft:  bl,
			BottomRight: br,
		}
	}

	ans := build(0, 0, n)
	return ans
}

func main() {
	construct([][]int{{1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 1, 1, 1, 1}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}, {1, 1, 1, 1, 0, 0, 0, 0}})
	construct([][]int{{0, 1}, {1, 0}})
	construct([][]int{{1, 1}, {1, 1}})
}
