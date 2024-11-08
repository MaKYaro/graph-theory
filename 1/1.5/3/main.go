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

		graph[v2] = append(graph[v2], v1)
	}

	result := dfs(graph, n)

	sort.Slice(result, func(i, j int) bool { return result[i] < result[j] })

	for i := 0; i < len(result); i++ {
		fmt.Printf("%d ", result[i])
	}
	fmt.Println()
}

func dfs(graph map[int][]int, n int) []int {
	states := make([]int, n+1)
	states[1] = VISITING

	result := dfsVisit(graph, 1, states)

	return result
}

func dfsVisit(graph map[int][]int, v int, states []int) []int {
	result := make([]int, 0, 1)
	stack := make([]int, 0, 1)
	stack = append(stack, v)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		nei := firstUnvisitedNeighbor(graph, curr, states)
		if nei == NIL {
			stack = stack[:len(stack)-1]
			states[curr] = VISITED
			result = append(result, curr)
		} else {
			states[nei] = VISITING
			stack = append(stack, nei)
		}
	}

	return result
}

func firstUnvisitedNeighbor(graph map[int][]int, v int, states []int) int {
	for _, nei := range graph[v] {
		if states[nei] == UNVISITED {
			//graph[v] = graph[v][1:]
			//fmt.Println(graph[v])
			return nei
		}
	}
	return NIL
}
