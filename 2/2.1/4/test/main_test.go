package test

import "testing"

func TestBfs(t *testing.T) {
	start := NewCell(0, 0)
	finish := NewCell(0, 2)

	got := bfs(start, finish)

	if got != 1 {
		t.Error("failed")
	}
}

const (
	UNVISITED = 0
	VISITING  = 1
	VISITED   = 2
	NIL       = -1
)

// func convert(coords string) Cell {
// 	letter, digit := coords[0], coords[1]

// 	x := int(letter - 'a')
// 	y := int(digit - '1')

// 	return NewCell(x, y)
// }

type Cell struct {
	x, y int
}

func (c Cell) Equal(t Cell) bool {
	return c.x == t.x && c.y == t.y
}

func (c Cell) Valid(n int) bool {
	return c.x >= 0 && c.x < n && c.y >= 0 && c.y < n
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

func bfs(start, finish Cell) int {
	if start.Equal(finish) {
		return 0
	}
	states := make(map[Cell]int, 64)
	states[start] = VISITING

	q := NewQueue(64)
	q.Push(start)

	distance := 0

	for q.Size() > 0 {
		l := q.Size()
		distance++

		for l > 0 {
			curr := q.Pop()

			for _, nei := range findUnvisitedNeighbours(states, curr, 8) {
				if nei.Equal(finish) && distance%2 == 0 {
					return distance / 2
				}
				if nei.Equal(finish) {
					return -1
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

func findUnvisitedNeighbours(graph map[Cell]int, c Cell, n int) []Cell {
	result := make([]Cell, 0, 8)
	x, y := c.x, c.y

	nei := NewCell(x+1, y+2)
	if nei.Valid(n) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x+2, y+1)
	if nei.Valid(n) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x+2, y-1)
	if nei.Valid(n) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x+1, y-2)
	if nei.Valid(n) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x-1, y-2)
	if nei.Valid(n) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x-2, y-1)
	if nei.Valid(n) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x-2, y+1)
	if nei.Valid(n) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}
	nei = NewCell(x-1, y+2)
	if nei.Valid(n) && graph[nei] == UNVISITED {
		result = append(result, nei)
	}

	return result
}
