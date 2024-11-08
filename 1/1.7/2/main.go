package main

import "fmt"

const (
	UNVISITED = 0
	VISITING  = 1
	VISITED   = 2
	NIL       = -1
)

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	graph := make(map[int][]int, n)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < m; i++ {
		var v1, v2 int
		fmt.Scan(&v1, &v2)

		graph[v1] = append(graph[v1], v2)
		graph[v2] = append(graph[v2], v1)
	}

	if isTree(graph, n) {
		fmt.Println("YES")
	} else {
		fmt.Println("NO")
	}

}

func isTree(graph map[int][]int, n int) bool {
	states := make([]int, n+1)
	parents := make([]int, n+1)
	compNum := 0

	for v := range graph {
		if states[v] == UNVISITED {
			compNum++
			if hasCicle(graph, states, parents, v, n) {
				return false
			}
		}
	}

	if compNum > 1 {
		return false
	}

	return true
}

func hasCicle(graph map[int][]int, states, parents []int, v, n int) bool {
	states[v] = VISITING
	stack := make([]int, 0, n)
	stack = append(stack, v)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		if hasVisitingNei(graph, states, parents, curr) {
			return true
		}
		nei := firstUnvisitedNei(graph, states, curr)
		if nei == NIL {
			stack = stack[:len(stack)-1]
			states[curr] = VISITED
		} else {
			states[nei] = VISITING
			stack = append(stack, nei)
			parents[nei] = curr
		}
	}

	return false
}

func firstUnvisitedNei(graph map[int][]int, states []int, v int) int {
	for _, nei := range graph[v] {
		if states[nei] == UNVISITED {
			return nei
		}
	}

	return NIL
}

func hasVisitingNei(graph map[int][]int, states, parents []int, v int) bool {
	for _, nei := range graph[v] {
		if states[nei] == VISITING && parents[v] != nei {
			return true
		}
	}

	return false
}
