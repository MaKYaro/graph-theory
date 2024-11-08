package main

import "fmt"

const (
	UNVISITED = 0
	VISITING  = 1
	VISITED   = 2
	NIL       = -1
)

func main() {
	var n int
	fmt.Scan(&n)

	graph := make(map[int][]int, n)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, 0)
	}

	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			var v int
			fmt.Scan(&v)

			if v == 1 {
				graph[i] = append(graph[i], j)
			}
		}
	}

	if dfs(graph, n) {
		fmt.Println(1)
	} else {
		fmt.Println(0)
	}
	//fmt.Println(graph)

}

func dfs(graph map[int][]int, n int) bool {
	states := make([]int, n+1)

	for v := range graph {
		//fmt.Println(states)
		if states[v] == UNVISITED {
			if dfsVisit(graph, states, v, n) {
				//fmt.Println(states)
				return true
			}
		}
	}

	return false
}

func dfsVisit(graph map[int][]int, states []int, v, n int) bool {
	states[v] = VISITING
	stack := make([]int, 0, n)
	stack = append(stack, v)

	for len(stack) > 0 {
		curr := stack[len(stack)-1]
		if hasVisitingNei(graph, states, curr) {
			return true
		}
		nei := firstUnvisitedNei(graph, states, curr)
		if nei == NIL {
			stack = stack[:len(stack)-1]
			states[curr] = VISITED
		} else {
			states[nei] = VISITING
			stack = append(stack, nei)
		}

	}
	return false
}

func hasVisitingNei(graph map[int][]int, states []int, v int) bool {
	for _, nei := range graph[v] {
		if states[nei] == VISITING {
			//fmt.Printf("%d has visiting nei: %d\n", v, nei)
			return true
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
