package main

//
// import (
// 	"strconv"
// 	"strings"
// )
//
// /*
//  Input: (2 -> 4 -> 3) + (5 -> 6 -> 4)
//  Output: 7 -> 0 -> 8
// */
//
// // Definition for singly-linked list.
// type ListNode struct {
// 	Val  int
// 	Next *ListNode
// }
//
// func addTwoNumbers(l1 *ListNode, l2 *ListNode) *ListNode {
// 	head := &ListNode{}
//
// 	n1Digits := make([]int, 0)
// 	n2Digits := make([]int, 0)
//
// 	for l1 != nil {
// 		n1Digits = append(n1Digits, l1.Val)
// 		l1 = l1.Next
// 	}
//
// 	for l2 != nil {
// 		n2Digits = append(n2Digits, l2.Val)
// 		l2 = l2.Next
// 	}
//
// 	n1, n2 := digitsToNum(n1Digits), digitsToNum(n2Digits)
//
// 	s := strconv.Itoa(n1 + n2)
//
// 	chars := strings.Split(s, "")
//
// 	curr := head
// 	head.Next = curr
// 	for i := len(chars) - 1; i >= 0; i-- {
// 		val, _ := strconv.Atoi(chars[i])
// 		curr.Val = val
// 		newNode := &ListNode{}
// 		if i > 0 {
// 			curr.Next = newNode
// 			curr = curr.Next
// 		}
// 	}
//
// 	return head
// }
//
// func digitsToNum(digits []int) int {
// 	var num int
// 	multiplier := 1
// 	for _, item := range digits {
// 		num += item * multiplier
// 		multiplier *= 10
// 	}
// 	return num
// }
