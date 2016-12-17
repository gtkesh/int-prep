package data_structures

import "fmt"

type BST struct {
	root *node
	size int
}

type node struct {
	left  *node
	right *node
	val   int
}

func (bst *BST) String() string {
	if bst.root == nil {
		return ""
	}

	return fmt.Sprint(bst.root.left, bst.root.val, bst.root.right)
}

func (bst *BST) insert(root *node, val int) *node {
	if root == nil {
		root = &node{val: val}
		bst.size += 1
		return root
	}

	switch {
	case val < root.val:
		root.left = bst.insert(root.left, val)
	case val >= root.val:
		root.right = bst.insert(root.right, val)
	}

	return root
}

// Find: O(n).
func (bst *BST) find(root *node, val int) bool {
	if root == nil {
		return false
	}
	return root.val == val || bst.find(root.left, val) || bst.find(root.right, val)
}

// Search: does the same thing as Find but faster - O(logn).
func (bst *BST) search(root *node, val int) bool {
	if root == nil {
		return false
	}

	switch {
	case val == root.val:
		return true
	case val < root.val:
		return bst.search(root.left, val)
	case val > root.val:
		return bst.search(root.right, val)
	default:
		return false
	}
}

// Sum: O(n).
func (bst *BST) sum(root *node) int {
	if root == nil {
		return 0
	}
	return root.val + bst.sum(root.left) + bst.sum(root.right)
}

func (bst *BST) traversePreorder(root *node) {
	if root == nil {
		return
	}

	fmt.Printf("%v ", root.val)

	if root.left != nil {
		bst.traversePreorder(root.left)
	}

	if root.right != nil {
		bst.traversePreorder(root.right)
	}
}

func (bst *BST) traverseInorder(root *node) {
	if root == nil {
		return
	}

	if root.left != nil {
		bst.traverseInorder(root.left)
	}

	fmt.Printf("%v ", root.val)

	if root.right != nil {
		bst.traverseInorder(root.right)
	}

}

func (bst *BST) traversePostorder(root *node) {
	if root == nil {
		return
	}

	if root.left != nil {
		bst.traversePostorder(root.left)
	}

	if root.right != nil {
		bst.traversePostorder(root.right)
	}

	fmt.Printf("%v ", root.val)

}

// TODO:
func (bst *BST) remove(n *node) error {
	return nil
}

// TODO:
func (bst *BST) height(n *node) int {
	return 0
}
