package main

func minPathSum(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])

	return minCost(grid, m-1, n-1)
}

func minCost(grid [][]int, i, j int) int {
	if i < 0 || j < 0 {
		MaxInt64 := 1<<63 - 1
		return MaxInt64
	}

	if i == 0 && j == 0 {
		return grid[i][j]
	}

	return grid[i][j] + min(minCost(grid, i-1, j), minCost(grid, i, j-1))
}

func minPathSumMemoize(grid [][]int) int {
	m := len(grid)
	n := len(grid[0])
	memo := make(map[string]int)
	return minCostMemoize(grid, m-1, n-1, memo)
}

func minCostMemoize(grid [][]int, i, j int, memo map[string]int) int {
	if i < 0 || j < 0 {
		MaxInt64 := 1<<63 - 1
		return MaxInt64
	}

	if i == 0 && j == 0 {
		return grid[i][j]
	}

	if val, ok := memo[string(i)+"-"+string(j)]; ok {
		return val
	}

	cost := grid[i][j] + min(minCostMemoize(grid, i-1, j, memo), minCostMemoize(grid, i, j-1, memo))
	memo[string(i)+"-"+string(j)] = cost

	return cost
}

func min(a, b int) int {
	if a <= b {
		return a
	}
	return b
}
