package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	graph := make(map[int][]int, 0)
	for i := 1; i <= n; i++ {
		graph[i] = make([]int, 0)
	}

	for i := 0; i < m; i++ {
		var v1, v2 int
		fmt.Scan(&v1, &v2)

		graph[v1] = append(graph[v1], v2)
		graph[v2] = append(graph[v2], v1)
	}
}
