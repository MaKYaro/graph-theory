package main

import (
	"fmt"
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	table := make([][]int, n)
	for i := 0; i < n; i++ {
		table[i] = make([]int, m)
		for j := 0; j < m; j++ {
			var val int
			fmt.Scan(&val)
			table[i][j] = val
		}
	}

	result := bfs(table, n, m)

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			fmt.Printf("%d ", result[i][j])
		}
		fmt.Println()
	}
}

const (
	UNVISITED = 0
	VISITING  = 1
	VISITED   = 2
	NIL       = -1
)

func Abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

type Cell struct {
	x, y int
}

func (c Cell) Valid(n, m int) bool {
	return c.x >= 0 && c.x <= n-1 && c.y >= 0 && c.y <= m-1
}

func (c1 Cell) Distance(c2 Cell) int {
	return Abs(c1.x-c2.x) + Abs(c1.y-c2.y)
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

func bfs(table [][]int, n, m int) [][]int {
	result := make([][]int, n)
	for i := 0; i < n; i++ {
		result[i] = make([]int, m)
	}

	for i := 0; i < n; i++ {
		for j := 0; j < m; j++ {
			c := NewCell(i, j)
			distance := findDistance(table, c, n, m)
			result[i][j] = distance
		}
	}

	return result
}

func findDistance(table [][]int, c Cell, n, m int) int {
	if table[c.x][c.y] == 1 {
		return 0
	}

	states := make(map[Cell]int, n*m)
	states[c] = VISITING

	q := NewQueue(n * m)
	q.Push(c)

	for q.Size() > 0 {
		l := q.Size()
		//fmt.Println(q)

		for l > 0 {
			// fmt.Println(q)
			curr := q.Pop()
			//fmt.Println(curr)
			//fmt.Println(states)

			for _, nei := range findUnvisitedNeighbours(states, curr, n, m) {
				//fmt.Println(findUnvisitedNeighbours(states, c, n, m))
				if table[nei.x][nei.y] == 1 {
					//fmt.Println(nei)
					return c.Distance(nei)
				}
				q.Push(nei)
				states[nei] = VISITING
			}
			states[curr] = VISITED
			l--
		}
	}
	return 0
}

func findUnvisitedNeighbours(states map[Cell]int, c Cell, n, m int) []Cell {
	result := make([]Cell, 0, 8)
	x, y := c.x, c.y

	nei := NewCell(x, y+1)
	if nei.Valid(n, m) && states[nei] == UNVISITED {
		result = append(result, nei)
	}
	// nei = NewCell(x+1, y+1)
	// if nei.Valid(n, m) && states[nei] == UNVISITED {
	// 	result = append(result, nei)
	// }
	nei = NewCell(x+1, y)
	if nei.Valid(n, m) && states[nei] == UNVISITED {
		result = append(result, nei)
	}
	// nei = NewCell(x+1, y-1)
	// if nei.Valid(n, m) && states[nei] == UNVISITED {
	// 	result = append(result, nei)
	// }
	nei = NewCell(x, y-1)
	if nei.Valid(n, m) && states[nei] == UNVISITED {
		result = append(result, nei)
	}
	// nei = NewCell(x-1, y-1)
	// if nei.Valid(n, m) && states[nei] == UNVISITED {
	// 	result = append(result, nei)
	// }
	nei = NewCell(x-1, y)
	if nei.Valid(n, m) && states[nei] == UNVISITED {
		result = append(result, nei)
	}
	// nei = NewCell(x-1, y+1)
	// if nei.Valid(n, m) && states[nei] == UNVISITED {
	// 	result = append(result, nei)
	// }

	return result
}
