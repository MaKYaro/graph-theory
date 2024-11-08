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

	components := findComponents(graph)

	sort.Slice(components, func(i, j int) bool { return components[i][0] < components[j][0] })

	fmt.Printf("%d\n", len(components))
	for i := 0; i < len(components); i++ {
		compLen := len(components[i])
		fmt.Printf("%d\n", compLen)

		for j := 0; j < compLen; j++ {
			fmt.Printf("%d ", components[i][j])
		}
		fmt.Println()
	}
}

func findComponents(graph map[int][]int) [][]int {
	components := dfs(graph)

	return components
}

func dfs(graph map[int][]int) [][]int {
	n := len(graph)
	states := make([]int, n+1)
	// for i := 0; i < n; i++ {
	// 	states[i] = UNVISITED
	// }

	components := make([][]int, 0, 1)
	stack := make([]int, 0, len(graph))

	for v := range graph {
		if states[v] == UNVISITED {
			component := dfsVisit(graph, v, states, stack)
			components = append(components, component)
		}
	}

	return components
}

func dfsVisit(graph map[int][]int, v int, states []int, stack []int) []int {
	states[v] = VISITING

	stack = append(stack, v)

	component := make([]int, 0, 1)

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

	sort.Slice(component, func(j, k int) bool { return component[j] < component[k] })

	return component
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
