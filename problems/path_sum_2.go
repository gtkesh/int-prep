package main

import "fmt"

/*

Given a binary tree and a sum, find all root-to-leaf paths where each path's sum equals the given sum.

For example:
Given the below binary tree and sum = 22,
              5
             / \
            4   8
           /   / \
          11  13  4
         /  \    / \
        7    2  5   1
return
[
   [5,4,11,2],
   [5,8,4,5]
]

*/

// TODO:(Debug)
func pathSum(root *TreeNode, sum int) [][]int {
	paths := make([][]int, 0)

	if root != nil {
		path := make([]int, 0)
		path = append(path, root.Val)
		traverse(root, sum-root.Val, path, paths)
	}

	return paths
}

func traverse(node *TreeNode, sum int, path []int, paths [][]int) {
	fmt.Println(path)
	fmt.Println(paths)
	if node.Right != nil && node.Left != nil && sum == 0 {
		p := make([]int, 0)
		for _, item := range path {
			p = append(p, item)
		}
		paths = append(paths, p)
	}

	if node.Left != nil {
		path = append(path, node.Left.Val)
		traverse(node.Left, sum-node.Left.Val, path, paths)
		// pop
		pathCopy := path
		pathCopy = pathCopy[:len(pathCopy)-1]
		path = pathCopy
	}

	if node.Right != nil {
		path = append(path, node.Right.Val)
		traverse(node.Right, sum-node.Right.Val, path, paths)
		// pop
		pathCopy := path
		pathCopy = pathCopy[:len(pathCopy)-1]
		path = pathCopy
	}
}
