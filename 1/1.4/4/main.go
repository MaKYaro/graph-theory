package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
	}

	for i := 0; i < m; i++ {
		var v1, v2 int
		fmt.Scan(&v1, &v2)

		graph[v1-1][v2-1]++
		graph[v2-1][v1-1]++
	}

	result := 0
	for i := 0; i < n; i++ {
		for j := i; j < n; j++ {
			if graph[i][j] > 1 {
				result++
			}
		}
	}
	fmt.Println(result)
}
