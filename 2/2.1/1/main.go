package main

import (
	"fmt"
)

func main() {
	var n int
	fmt.Scan(&n)

	// create graph
	graph := make(map[int][]int, n)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, 0)
	}

	// read graph
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			var val int
			fmt.Scan(&val)

			if val == 1 {
				graph[i] = append(graph[i], j)
			}
		}
	}

	var v1, v2 int
	fmt.Scan(&v1, &v2)

	fmt.Println(Bfs(graph, n, v1, v2))
}

const (
	UNVISITED = 0
	VISITING  = 1
	VISITED   = 2
	NIL       = -1
)

type Queue struct {
	items []int
}

func NewQueue(size int) *Queue {
	items := make([]int, 0, size)
	return &Queue{items: items}
}

func (q *Queue) Push(elem int) {
	q.items = append(q.items, elem)
}

func (q *Queue) Size() int {
	return len(q.items)
}

func (q *Queue) First() int {
	return q.items[0]
}

func (q *Queue) Pop() int {
	first := q.First()
	if q.Size() == 1 {
		q.items = q.items[:0]
	} else {
		q.items = q.items[1:]
	}

	return first
}

func Bfs(graph map[int][]int, n, s, f int) int {
	if s == f {
		return 0
	}
	states := make([]int, n+1)
	states[s] = VISITING
	distance := 0

	q := NewQueue(n)
	q.Push(s)

	for q.Size() > 0 {
		distance++

		l := q.Size()
		for l > 0 {
			curr := q.Pop()

			for _, nei := range graph[curr] {
				if nei == f {
					return distance
				}
				if states[nei] == UNVISITED {
					states[nei] = VISITING
					q.Push(nei)
				}
			}
			states[curr] = VISITED
			l--
		}
	}

	return -1
}
