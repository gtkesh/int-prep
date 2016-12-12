package main

import (
	"fmt"
	"testing"
)

func TestSize(t *testing.T) {
	// Create a BST.
	bst := new(BST)

	// Insert 5 items: 0,1,2,3,4.
	items := 5
	for i := 0; i < items; i++ {
		bst.root = bst.insert(bst.root, i)
	}

	// Assert the size.
	if bst.size != items {
		t.Error(
			"expected", items,
			"got", bst.size,
		)
	}

	bst = new(BST)
	// Insert 5 items: 4,3,2,1,0.
	for i := items - 1; i >= 0; i-- {
		bst.root = bst.insert(bst.root, i)
	}

	// Assert the size.
	if bst.size != items {
		t.Error(
			"expected", items,
			"got", bst.size,
		)
	}
}

func TestSearch(t *testing.T) {
	bst := new(BST)

	// Insert 5 items: 0,1,2,3,4
	items := 5
	for i := 0; i < items; i++ {
		bst.root = bst.insert(bst.root, i)
	}

	// Search for 5 => must return false.
	if bst.search(bst.root, 5) {
		t.Error(
			"expected", false,
			"got", true,
		)
	}

	// Search for -1 => must return false.
	if bst.search(bst.root, -1) {
		t.Error(
			"expected", false,
			"got", true,
		)
	}

	// Search for 0,1,2,3,4 => must return true for all.
	for i := 0; i < items; i++ {
		exists := bst.search(bst.root, i)
		if !exists {
			t.Error(
				"expected", true,
				"got", false,
			)
		}
	}
}

func TestFind(t *testing.T) {
	bst := new(BST)

	// Insert 5 items: 0,1,2,3,4
	items := 5
	for i := 0; i < items; i++ {
		bst.root = bst.insert(bst.root, i)
	}

	// Search for 5 => must return false.
	if bst.find(bst.root, 5) {
		t.Error(
			"expected", false,
			"got", true,
		)
	}

	// Search for -1 => must return false.
	if bst.find(bst.root, -1) {
		t.Error(
			"expected", false,
			"got", true,
		)
	}

	// Search for 0,1,2,3,4 => must return true for all.
	for i := 0; i < items; i++ {
		exists := bst.find(bst.root, i)
		if !exists {
			t.Error(
				"expected", true,
				"got", false,
			)
		}
	}

	// Empty tree.
	exists := bst.find(nil, 10)
	if exists {
		t.Error(
			"expected", false,
			"got", true,
		)
	}
}

func TestSum(t *testing.T) {
	bst := new(BST)

	// Insert 5 items: 0,1,2,3,4.
	// Sum = 0+1+2+3+4 = 10
	expectedSum := 0
	items := 5
	for i := 0; i < items; i++ {
		bst.root = bst.insert(bst.root, i)
		expectedSum += i
	}

	// Assert the sum.
	sum := bst.sum(bst.root)
	if sum != expectedSum {
		t.Error(
			"expected", expectedSum,
			"got", sum,
		)
	}
}

func TestTraversal(t *testing.T) {
	bst := new(BST)
	// Tree structure
	/*
	        40
	       /  \
	     25    78
	    /  \
	   10  32

	*/

	bst.root = bst.insert(bst.root, 40)
	bst.root = bst.insert(bst.root, 25)
	bst.root = bst.insert(bst.root, 78)
	bst.root = bst.insert(bst.root, 10)
	bst.root = bst.insert(bst.root, 32)

	// Print Preorder
	bst.traversePreorder(bst.root)
	fmt.Println()
	// Print Inorder
	bst.traverseInorder(bst.root)
	fmt.Println()
	// Print Postorder
	bst.traversePostorder(bst.root)
	fmt.Println()

}
