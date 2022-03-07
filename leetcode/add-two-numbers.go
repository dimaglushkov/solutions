package main

// source: https://leetcode.com/problems/add-two-numbers/

/*type ListNode struct {
	Val  int
	Next *ListNode
}
func buildList(num []int) *ListNode {
	var prevNode, iNode *ListNode
	for _, i := range num {
		iNode = &ListNode{Val: i, Next: prevNode}
		prevNode = iNode
	}
	return iNode
}
*/
// Basically implementing 1st grade school algorithm to manually sum to numbers
func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
	var curNode *ListNode
	var nextNode = &ListNode{}
	res := nextNode

	for l1 != nil || l2 != nil {
		curNode = nextNode
		nextNode = &ListNode{}

		if l1 != nil {
			curNode.Val += l1.Val
			l1 = l1.Next
		}
		if l2 != nil {
			curNode.Val += l2.Val
			l2 = l2.Next
		}
		if curNode.Val > 9 {
			nextNode.Val = 1
			curNode.Val %= 10
		}
		curNode.Next = nextNode
	}

	if nextNode.Val == 0 {
		curNode.Next = nil
	}
	return res
}

func main() {
	l1 := buildList([]int{9, 9, 9, 9, 9, 9, 9})
	l2 := buildList([]int{9, 9, 9, 9})
	addTwoNumbers(l1, l2)
}
