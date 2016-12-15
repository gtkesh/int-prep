package main

type ListNode struct {
	Val  int
	Next *ListNode
}

func mergeTwoLists(l1 *ListNode, l2 *ListNode) *ListNode {
	head := new(ListNode)
	curr := head
	for l1 != nil && l2 != nil {
		if l1.Val <= l2.Val {
			curr.Next = &ListNode{Val: l1.Val}
			l1 = l1.Next
		} else {
			curr.Next = &ListNode{Val: l2.Val}
			l2 = l2.Next
		}
		curr = curr.Next
	}

	return head.Next
}
