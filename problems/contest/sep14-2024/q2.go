package sep142024

/*
Q2. Find a Safe Walk Through a Grid
Solved
Medium
5 pt.
You are given an m x n binary matrix grid and an integer health.

You start on the upper-left corner (0, 0) and would like to get to the lower-right corner (m - 1, n - 1).

You can move up, down, left, or right from one cell to another adjacent cell as long as your health remains positive.

Cells (i, j) with grid[i][j] = 1 are considered unsafe and reduce your health by 1.

Return true if you can reach the final cell with a health value of 1 or more, and false otherwise.



Example 1:

Input: grid = [[0,1,0,0,0],[0,1,0,1,0],[0,0,0,1,0]], health = 1

Output: true

Explanation:

The final cell can be reached safely by walking along the gray cells below.


Example 2:

Input: grid = [[0,1,1,0,0,0],[1,0,1,0,0,0],[0,1,1,1,0,1],[0,0,1,0,1,0]], health = 3

Output: false

Explanation:

A minimum of 4 health points is needed to reach the final cell safely.


Example 3:

Input: grid = [[1,1,1],[1,0,1],[1,1,1]], health = 5

Output: true

Explanation:

The final cell can be reached safely by walking along the gray cells below.



Any path that does not go through the cell (1, 1) is unsafe since your health will drop to 0 when reaching the final cell.



Constraints:

m == grid.length
n == grid[i].length
1 <= m, n <= 50
2 <= m * n
1 <= health <= m + n
grid[i][j] is either 0 or 1.
*/

var directions = [][]int{
	{0, 1},
	{1, 0},
	{0, -1},
	{-1, 0},
}

func findSafeWalk(grid [][]int, health int) bool {
	m, n := len(grid), len(grid[0])

	q := make([]struct {
		pos []int
		h   int
	}, 1)
	q[0].pos = []int{0, 0}
	q[0].h = health - grid[0][0]

	v := make([][]int, m)
	for i := range v {
		v[i] = make([]int, n)
	}
	v[0][0] = q[0].h

	for len(q) > 0 {
		var p struct {
			pos []int
			h   int
		}
		p, q = q[0], q[1:]

		x, y := p.pos[0], p.pos[1]
		if x == m-1 && y == n-1 && p.h > 0 {
			return true
		}

		for _, dir := range directions {
			nx, ny := x+dir[0], y+dir[1]
			if nx >= 0 && nx < m && ny >= 0 && ny < n {
				nh := p.h - grid[nx][ny]
				if nh > 0 && nh > v[nx][ny] {
					v[nx][ny] = nh
					q = append(q, struct {
						pos []int
						h   int
					}{
						pos: []int{nx, ny},
						h:   nh,
					})
				}
			}
		}
	}

	return false
}
