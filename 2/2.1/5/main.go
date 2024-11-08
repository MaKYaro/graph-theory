package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	field := make([][]int, n)
	for i := 0; i < n; i++ {
		field[i] = make([]int, m)
		for j := 0; j < m; j++ {
			var v int
			fmt.Scan(&v)
			field[i][j] = v
		}
	}

	inclines := bfs(field, n, m)
	fmt.Println(inclines)
}

const (
	UNVISITED = 0
	VISITING  = 1
	VISITED   = 2
)

type Cell struct {
	x, y int
}

func (c Cell) Valid(n, m int) bool {
	return c.x >= 0 && c.x < n && c.y >= 0 && c.y < m
}

func NewCell(x, y int) Cell {
	return Cell{x, y}
}

type Queue struct {
	items []Cell
}

func NewQueue(size int) *Queue {
	items := make([]Cell, 0, size)
	return &Queue{items: items}
}

func (q *Queue) Push(item Cell) {
	q.items = append(q.items, item)
}

func (q *Queue) First() Cell {
	return q.items[0]
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) Pop() Cell {
	first := q.First()
	if q.Size() == 1 {
		q.items = q.items[:0]
	} else {
		q.items = q.items[1:]
	}

	return first
}

func bfs(field [][]int, n, m int) int {
	start := NewCell(0, 0)
	states := make(map[Cell]int, n)
	states[start] = VISITED

	q := NewQueue(n * m)
	q.Push(start)

	inclines := 0

	for q.Size() > 0 {
		l := q.Size()
		inclines++

		for l > 0 {
			curr := q.Pop()
			for _, nei := range findUnvisitedNeighbours(field, states, n, m, curr) {
				if pathHasHole(field, curr, nei) {
					return inclines
				}
				states[nei] = VISITING
				q.Push(nei)
			}
			states[curr] = VISITED
			l--
		}
	}

	return inclines
}

func pathHasHole(table [][]int, start, finish Cell) bool {
	if start.x < finish.x {
		x1, x2 := start.x, finish.x
		y := start.y
		for x1 < x2 {
			x1++
			if table[y][x1] == 2 {
				return true
			}
		}
	}
	if start.x > finish.x {
		x1, x2 := start.x, finish.x
		y := start.y
		for x1 > x2 {
			x1--
			if table[y][x1] == 2 {
				return true
			}
		}
	}
	if start.y < finish.y {
		y1, y2 := start.y, finish.y
		x := start.x
		for y1 < y2 {
			y1++
			if table[y1][x] == 2 {
				return true
			}
		}
	}
	if start.y > finish.y {
		y1, y2 := start.y, finish.y
		x := start.x
		for y1 > y2 {
			y1--
			if table[y1][x] == 2 {
				return true
			}
		}
	}

	return false
}

func findUnvisitedNeighbours(
	table [][]int,
	states map[Cell]int,
	n, m int, c Cell,
) []Cell {
	result := make([]Cell, 0, 4)

	// find top neighbour
	x, y := c.x, c.y
	for validCoords(table, n, m, x, y) {
		y--
	}
	y++
	result = addUnvisitedNei(states, result, x, y)

	// find bottom neighbour
	x, y = c.x, c.y
	for validCoords(table, n, m, x, y) {
		y++
	}
	y--
	result = addUnvisitedNei(states, result, x, y)

	// find left neighbour
	x, y = c.x, c.y
	for validCoords(table, n, m, x, y) {
		x--
	}
	x++
	result = addUnvisitedNei(states, result, x, y)

	// find right neighbour
	x, y = c.x, c.y
	for validCoords(table, n, m, x, y) {
		x++
	}
	x--
	result = addUnvisitedNei(states, result, x, y)

	return result
}

func addUnvisitedNei(states map[Cell]int, result []Cell, x, y int) []Cell {
	nei := NewCell(x, y)
	if states[nei] == UNVISITED {
		result = append(result, NewCell(x, y))
	}
	return result
}

func validCoords(table [][]int, n, m, x, y int) bool {
	validX := x >= 0 && x < m
	validY := y >= 0 && y < n
	var hasZero, hasTwo bool
	if validX && validY {
		hasZero = table[y][x] == 0
		hasTwo = table[y][x] == 2
	}
	//validValue := table[x][y] == 0 || table[x][y] == 2

	return validX && validY && (hasZero || hasTwo)
}
