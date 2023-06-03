package main

//
//import (
//	"fmt"
//)
//
//func max(x, y int) int {
//	if x > y {
//		return x
//	}
//	return y
//}
//
//type Tree struct {
//	w        int
//	children []*Tree
//}
//
//func buildTree(rootId int, edges, weights []int) *Tree {
//	root := &Tree{
//		w:        weights[rootId],
//		children: nil,
//	}
//	for i := range edges {
//		if edges[i] != rootId {
//			continue
//		}
//		root.children = append(root.children, buildTree(i, edges, weights))
//	}
//
//	return root
//}
//
//func dfs(curSum int, root *Tree) int {
//	curSum += root.w
//	maxSum := curSum
//	for _, ch := range root.children {
//		maxSum = max(maxSum, dfs(curSum, ch))
//	}
//	return maxSum
//}
//
//func numOfMinutes(n int, headID int, manager []int, informTime []int) int {
//	return dfs(0, buildTree(headID, manager, informTime))
//}
//
//func main() {
//	testCases := []struct {
//		n          int
//		headID     int
//		manager    []int
//		informTime []int
//		want       int
//	}{
//		{
//			n:          6,
//			headID:     2,
//			manager:    []int{2, 2, -1, 2, 2, 2},
//			informTime: []int{0, 0, 1, 0, 0, 0},
//			want:       1,
//		},
//		{
//			n:          1,
//			headID:     0,
//			manager:    []int{-1},
//			informTime: []int{0},
//			want:       0,
//		},
//	}
//
//	successes := 0
//	for _, tc := range testCases {
//		x := numOfMinutes(tc.n, tc.headID, tc.manager, tc.informTime)
//		status := "ERROR"
//		if fmt.Sprint(x) == fmt.Sprint(tc.want) {
//			status = "OK"
//			successes++
//		}
//		fmt.Println(status, "	Expected: ", tc.want, " Actual: ", x)
//	}
//	if l := len(testCases); successes == len(testCases) {
//		fmt.Printf("===\nSUCCESS: %d of %d tests ended successfully\n", successes, l)
//	} else {
//		fmt.Printf("===\nFAIL: %d tests failed\n", l-successes)
//	}
//
//}
