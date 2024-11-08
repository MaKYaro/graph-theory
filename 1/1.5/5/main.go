package main

import (
	"fmt"
	"sort"
)

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
	for v := range graph {
		sort.Slice(graph[v], func(i, j int) bool { return graph[v][i] < graph[v][j] })
	}
	var u, v int
	fmt.Scan(&u, &v)

	result := dfs(graph, n, u, v)
	if len(result) == 0 {
		fmt.Println("no solution")
		return
	}
	for i := 0; i < len(result)-1; i++ {
		fmt.Printf("%d ", result[i])
	}
	fmt.Println(result[len(result)-1])
}

func dfs(graph map[int][]int, n int, u, v int) []int {
	states := make([]int, n+1)
	states[u] = VISITING

	result := dfsVisit(graph, states, u, v)
	return result
}

func dfsVisit(graph map[int][]int, states []int, u, v int) []int {
	//result := make([]int, 0)
	stack := make([]int, 0, len(graph))
	stack = append(stack, u)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		nei := firstUnvisitedNei(graph, states, curr)
		if nei == v {
			stack = append(stack, nei)
			return stack
		}
		if nei == NIL {
			stack = stack[:len(stack)-1]
			states[curr] = VISITED
		} else {
			states[nei] = VISITING
			stack = append(stack, nei)
		}
	}

	return []int{}
}

func firstUnvisitedNei(graph map[int][]int, states []int, v int) int {
	for _, nei := range graph[v] {
		if states[nei] == UNVISITED {
			return nei
		}
	}
	return NIL
}
