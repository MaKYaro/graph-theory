package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	table := make([][][]byte, n)
	for i := 0; i < n; i++ {
		//  add block
		table[i] = make([][]byte, n)
		for j := 0; j < n; j++ {
			// add string to a block
			//table[i][j] = make([]byte, n)
			var s string
			fmt.Scan(&s)
			table[i][j] = []byte(s)
		}
	}

	result := bfs(table, n)
	fmt.Println(result)
}

const (
	UNVISITED = 0
	VISITING  = 1
	VISITED   = 2

	STONE      = '#'
	EMPTY_CELL = '.'
	START      = 'S'
)

// Cell is an element of 3d matrix
type Cell struct {
	x int // block number
	y int // string number
	z int // symbol number
}

func NewCell(x, y, z int) Cell {
	return Cell{x, y, z}
}

func (c Cell) Valid(n int) bool {
	return c.x >= 0 && c.x < n && c.y >= 0 && c.y < n && c.z >= 0 && c.z < n
}

func (c Cell) Empty(table [][][]byte) bool {
	return table[c.x][c.y][c.z] == EMPTY_CELL
}

// Queue is a FIFO datatype
type Queue struct {
	items []Cell
}

func NewQueue(size int) *Queue {
	items := make([]Cell, 0, size)
	return &Queue{items: items}
}

func (q *Queue) Push(c Cell) {
	q.items = append(q.items, c)
}

func (q *Queue) First() Cell {
	first := q.items[0]
	return first
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

func bfs(table [][][]byte, n int) int {
	start := findStart(table, n)
	if start.x == 0 {
		return 0
	}
	states := make(map[Cell]int, n*n*n)
	states[start] = VISITING

	distance := 0

	q := NewQueue(n * n * n)
	q.Push(start)

	for q.Size() > 0 {
		distance++
		l := q.Size()

		for l > 0 {
			curr := q.Pop()
			for _, nei := range findValidNeighbours(table, states, n, curr) {
				if nei.x == 0 {
					return distance
				}
				states[nei] = VISITING
				q.Push(nei)
			}
			states[curr] = VISITED
			l--
		}
	}

	return distance
}

func findValidNeighbours(table [][][]byte, states map[Cell]int, n int, c Cell) []Cell {
	result := make([]Cell, 0, 6)
	x, y, z := c.x, c.y, c.z

	nei := NewCell(x+1, y, z)
	if nei.Valid(n) && nei.Empty(table) && states[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x-1, y, z)
	if nei.Valid(n) && nei.Empty(table) && states[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x, y+1, z)
	if nei.Valid(n) && nei.Empty(table) && states[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x, y-1, z)
	if nei.Valid(n) && nei.Empty(table) && states[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x, y, z+1)
	if nei.Valid(n) && nei.Empty(table) && states[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x, y, z-1)
	if nei.Valid(n) && nei.Empty(table) && states[nei] == UNVISITED {
		result = append(result, nei)
	}

	return result
}

func findStart(table [][][]byte, n int) Cell {
	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			for k := 0; k < n; k++ {
				if table[i][j][k] == START {
					return NewCell(i, j, k)
				}
			}
		}
	}
	return NewCell(0, 0, 0)
}
