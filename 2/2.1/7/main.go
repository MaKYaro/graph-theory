package main

import (
	"fmt"
)

func main() {
	var n, m, mX, mY, fleasAmount int
	fmt.Scan(&n, &m, &mX, &mY, &fleasAmount)

	feeder := NewCell(mX-1, mY-1)

	fleas := make([]Cell, fleasAmount)
	for i := 0; i < fleasAmount; i++ {
		var x, y int
		fmt.Scan(&x, &y)
		fleas[i] = NewCell(x-1, y-1)
	}

	result := 0
	for _, flea := range fleas {
		distance := bfs(n, m, flea, feeder)
		if distance == -1 {
			fmt.Println(-1)
			return
		}
		result += distance
	}

	fmt.Println(result)
}

const (
	UNVISITED = 0
	VISITING  = 1
	VISITED   = 2
)

type Cell struct {
	x int // row number
	y int // elem number
}

func NewCell(x, y int) Cell {
	return Cell{x, y}
}

func (c1 Cell) Equal(c2 Cell) bool {
	return c1.x == c2.x && c1.y == c2.y
}

func (c Cell) Valid(n, m int) bool {
	return c.x >= 0 && c.x < n && c.y >= 0 && c.y < m
}

type Queue struct {
	items []Cell
}

func NewQueue(size int) *Queue {
	items := make([]Cell, 0, size)
	return &Queue{items}
}

func (q *Queue) Push(item Cell) {
	q.items = append(q.items, item)
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) First() Cell {
	return q.items[0]
}

func (q *Queue) Pop() Cell {
	first := q.First()
	if q.Size() > 1 {
		q.items = q.items[1:]
	} else {
		q.items = q.items[:0]
	}
	return first
}

func bfs(n, m int, flea Cell, feeder Cell) int {
	if flea.Equal(feeder) {
		return 0
	}
	states := make(map[Cell]int, n*m)
	states[flea] = VISITING

	q := NewQueue(n * m)
	q.Push(flea)

	distance := 0
	for q.Size() > 0 {
		l := q.Size()
		distance++

		for l > 0 {
			curr := q.Pop()
			for _, nei := range findValidUnvisitedNeighbours(states, curr, n, m) {
				if nei.Equal(feeder) {
					return distance
				}
				states[nei] = VISITING
				q.Push(nei)
			}
			states[curr] = VISITED
			l--
		}
	}

	return -1
}

func findValidUnvisitedNeighbours(graph map[Cell]int, c Cell, n, m int) []Cell {
	result := make([]Cell, 0, 8)
	x, y := c.x, c.y

	nei := NewCell(x+1, y+2)
	if nei.Valid(n, m) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x+2, y+1)
	if nei.Valid(n, m) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x+2, y-1)
	if nei.Valid(n, m) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x+1, y-2)
	if nei.Valid(n, m) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x-1, y-2)
	if nei.Valid(n, m) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x-2, y-1)
	if nei.Valid(n, m) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x-2, y+1)
	if nei.Valid(n, m) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x-1, y+2)
	if nei.Valid(n, m) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}

	return result
}
