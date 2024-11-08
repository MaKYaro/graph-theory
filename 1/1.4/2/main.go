package main

import "fmt"

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
	result := countVerticesDegrees(n, graph)
	fmt.Println(result)
}

func countVerticesDegrees(n int, graph map[int][]int) int {
	result := 0
	degrees := make(map[int]struct{}, n)

	for vert := range graph {
		degree := len(graph[vert])
		if _, ok := degrees[degree]; !ok {
			result++
			degrees[degree] = struct{}{}
		}
	}

	if len(graph) < n {
		result++
	}

	return result
}
