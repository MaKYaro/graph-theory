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
	for i := 0; i < m; i++ {
		var v1, v2 int
		fmt.Scan(&v1, &v2)
		graph[v1] = append(graph[v1], v2)
		graph[v2] = append(graph[v2], v1)
	}

	component := findComponent(graph, n)
	sort.Slice(component, func(i, j int) bool { return component[i] < component[j] })

	fmt.Println(len(component))
	for i := 0; i < len(component)-1; i++ {
		fmt.Printf("%d ", component[i])
	}
	fmt.Printf("%d\n", component[len(component)-1])
}

func findComponent(graph map[int][]int, n int) []int {
	component := make([]int, 0, n)
	states := make([]int, n+1)
	states[1] = VISITING
	stack := make([]int, 0, len(graph))
	stack = append(stack, 1)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		nei := firstUnvisitedNeighbor(graph, curr, states)
		if nei == NIL {
			stack = stack[:len(stack)-1]
			states[curr] = VISITED
			component = append(component, curr)
		} else {
			states[nei] = VISITING
			stack = append(stack, nei)
		}
	}

	return component

}

func firstUnvisitedNeighbor(graph map[int][]int, v int, states []int) int {
	for _, nei := range graph[v] {
		if states[nei] == UNVISITED {
			return nei
		}
	}
	return NIL
}

// func findComponent(graph map[int][]int, n int) []int {
// 	component := make([]int, 0, n)
// 	states := make([]int, n+1)

// 	var dfsVisit func(v int)
// 	dfsVisit = func(v int) {
// 		component = append(component, v)
// 		states[v] = VISITING
// 		for _, nei := range graph[v] {
// 			if states[nei] == UNVISITED {
// 				dfsVisit(nei)
// 			}
// 		}
// 		states[v] = VISITED
// 	}

// 	dfsVisit(1)

// 	return component
// }

// func dfs(graph map[int][]int) {
// 	n := len(graph)
// 	states := make([]int, n)
// 	for i := 0; i < n; i++ {
// 		states[i] = UNVISITED
// 	}

// 	for v := range graph {
// 		if states[v] == UNVISITED {
// 			dfsVisit(graph, v, states)
// 		}
// 	}
// }

// func dfsVisit(graph map[int][]int, v int, states []int) {
// 	states[v] = VISITING
// 	for _, nei := range graph[v] {
// 		if states[nei] == UNVISITED {
// 			dfsVisit(graph, nei, states)
// 		}
// 	}
// 	states[v] = VISITED
// }
