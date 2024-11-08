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
		sort.Slice(
			graph[v],
			func(i, j int) bool {
				return graph[v][i] < graph[v][j]
			},
		)
	}
	//fmt.Println(graph)
	hasCycle, cycle := hasCycleInGraph(graph, n)
	if hasCycle {
		fmt.Println("YES")
		printCycle(cycle)
	} else {
		fmt.Println("NO")
	}
}

func printCycle(cycle []int) {
	fmt.Println(len(cycle))
	for i := 0; i < len(cycle)-1; i++ {
		fmt.Printf("%d ", cycle[i])
	}
	fmt.Println(cycle[len(cycle)-1])
}

func hasCycleInGraph(graph map[int][]int, n int) (bool, []int) {
	states := make([]int, n+1)
	parents := make([]int, n+1)

	for v := 1; v <= n; v++ {
		if states[v] == UNVISITED {
			hasCycle, cycle := hasCycleInComponent(graph, states, parents, v, n)
			if hasCycle {
				return true, cycle
			}
		}
	}
	return false, []int{}
}

func hasCycleInComponent(graph map[int][]int, states, parents []int, v, n int) (bool, []int) {
	states[v] = VISITING
	stack := make([]int, 0, n)
	stack = append(stack, v)

	for len(stack) > 0 {
		//fmt.Println(stack)
		//fmt.Println(states)
		//fmt.Println(parents)
		curr := stack[len(stack)-1]
		if hasVisitingNei(graph, states, parents, curr) {
			return true, stack
		}
		nei := firstUnvisitedNei(graph, states, curr)
		if nei == NIL {
			states[curr] = VISITED
			stack = stack[:len(stack)-1]
		} else {
			states[nei] = VISITING
			stack = append(stack, nei)
			parents[nei] = curr
		}
	}
	return false, []int{}
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
