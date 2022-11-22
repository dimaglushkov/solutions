package main

import (
	"fmt"

	. "github.com/dimaglushkov/solutions/ADS/list"
)

// source: https://leetcode.com/problems/palindrome-linked-list/

/**
 * Definition for singly-linked list.
 * type ListNode struct {
 *     Val int
 *     Next *ListNode
 * }
 */
func isPalindrome(head *ListNode) bool {
	l := 0
	for p := head; p != nil; p = p.Next {
		l++
	}

	if l == 0 || l == 1 {
		return true
	}

	c := make([]int, 10)
	for i, p := 0, head; p != nil; i, p = i+1, p.Next {
		if i >= l/2 {
			c[p.Val] ^= i
		} else {
			c[p.Val] ^= l - 1 - i
		}
	}

	nonzeros := 0
	for _, v := range c {
		if v != 0 {
			nonzeros += 1
		}
	}

	return nonzeros == 0 || nonzeros == 1

}
func main() {
	// Example 1
	var head1 = NewList([]int{1, 2, 2, 1})
	fmt.Println("Expected: true	Output: ", isPalindrome(head1))

	// Example 2
	var head2 *ListNode = NewList([]int{1, 2})
	fmt.Println("Expected: false	Output: ", isPalindrome(head2))

}
