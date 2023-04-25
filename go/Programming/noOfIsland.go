package main

import (
	"fmt"
)

// solution (class) -> function

type Solution struct {
}

func (s *Solution) noOfIsland(grid [][]string) int {

	count := 0
	n := len(grid)
	m := len(grid[0])
	for i := 0; i < len(grid); i++ {
		for j := 0; j < len(grid[0]); j++ {
			if grid[i][j] == "1" {
				s.dfs(&grid, n, m, i, j)
				count = count + 1
			}
		}
	}
	return count
}

func (s *Solution) dfs(grid *[][]string, n int, m int, i int, j int) {

	if (i >= n) || (j >= m) || (i < 0) || (j < 0) || (*grid)[i][j] == "0" {
		return
	}

	(*grid)[i][j] = "0"
	s.dfs(grid, n, m, i+1, j)
	s.dfs(grid, n, m, i-1, j)
	s.dfs(grid, n, m, i, j+1)
	s.dfs(grid, n, m, i, j-1)

	return
}

func main() {

	grid := [][]string{
		{"1", "1", "0", "0", "0"},
		{"1", "1", "0", "0", "0"},
		{"0", "0", "1", "0", "0"},
		{"0", "0", "0", "1", "1"},
	}

	s := Solution{}

	fmt.Printf("number of island:=>%d\n", s.noOfIsland(grid))

}
