package main

import (
	"bufio"
	"fmt"
	"os"
)

type State struct {
	x, y, dir, steps int
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	var n, m int
	fmt.Fscan(reader, &n, &m)

	var sx, sy, tx, ty int
	fmt.Fscan(reader, &sx, &sy, &tx, &ty)

	sx--
	sy--
	tx--
	ty--

	grid := make([]string, n)

	for i := 0; i < n; i++ {
		fmt.Fscan(reader, &grid[i])

	}

	visited := make([][][2]bool, n)
	for i := range visited {
		visited[i] = make([][2]bool, m)
	}

	queue := []State{}

	queue = append(queue, State{sx, sy, 0, 0})
	queue = append(queue, State{sx, sy, 1, 0})

	visited[sx][sy][0] = true
	visited[sx][sy][1] = true

	for len(queue) > 0 {
		curr := queue[0]
		queue = queue[1:]
		if curr.x == tx && curr.y == ty {
			fmt.Println(curr.steps)
			return
		}

		var dx, dy []int
		var nextDir int
		if curr.dir == 0 {
			dx = []int{1, -1}
			dy = []int{0, 0}
			nextDir = 1
		} else {
			dy = []int{1, -1}
			dx = []int{0, 0}
			nextDir = 0
		}

		for i := 0; i < 2; i++ {
			nx, ny := curr.x+dx[i], curr.y+dy[i]

			if nx >= 0 && nx < n && ny >= 0 && ny < m && grid[nx][ny] == '.' && !visited[nx][ny][nextDir] {
				visited[nx][ny][nextDir] = true
				queue = append(queue, State{nx, ny, nextDir, curr.steps + 1})
			}
		}
	}
	fmt.Println("-1")
}
