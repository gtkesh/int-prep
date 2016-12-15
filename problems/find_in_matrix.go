package main

/*

Input: M=NxN matrix, K
Output: boolean -> True if K is in M. False otherwise.

1 2 3 4
5 6 7 8
9 10 11 12
13 14 15 16

*/

func containsNaive(matrix [][]int, k int) bool {
	n := len(matrix)

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if matrix[i][j] == k {
				return true
			}
		}
	}
	return false
}

func containsFast(matrix [][]int, k int) bool {
	n := len(matrix)

	for row, col := n-1, 0; row >= 0 && col <= n-1; {
		switch {
		case k == matrix[row][col]:
			return true
		case k > matrix[row][col]:
			col++
		case k < matrix[row][col]:
			row--
		}
	}

	return false
}
