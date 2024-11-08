package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	var x1, y1 int
	fmt.Scan(&x1, &y1)
	start := NewCell(x1, y1)

	var x2, y2 int
	fmt.Scan(&x2, &y2)
	finish := NewCell(x2, y2)

	graph := make(map[Cell]int, n*n)

	result := bfs(graph, start, finish, n)

	fmt.Println(result)
}

const (
	UNVISITED = 0
	VISITING  = 1
	VISITED   = 2
	NIL       = -1
)

type Cell struct {
	x, y int
}

func NewCell(x, y int) Cell {
	return Cell{x, y}
}

func (c Cell) Equal(t Cell) bool {
	return c.x == t.x && c.y == t.y
}

func (c Cell) Valid(n int) bool {
	return c.x >= 1 && c.x <= n && c.y >= 1 && c.y <= n
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

func bfs(graph map[Cell]int, start Cell, finish Cell, n int) int {
	if start.Equal(finish) {
		return 0
	}
	graph[start] = VISITING

	q := NewQueue(n * n)
	q.Push(start)

	distance := 0

	for q.Size() > 0 {
		distance++
		l := q.Size()

		for l > 0 {
			curr := q.Pop()
			for _, nei := range findUnvisitedNeighbours(graph, curr, n) {
				if nei.Equal(finish) {
					return distance
				}
				q.Push(nei)
			}
			graph[curr] = VISITED
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
