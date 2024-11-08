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
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			if i == j {
				continue
			}
			if graph[i][j] == graph[j][i] {
				fmt.Println("NO")
				return
			}
		}
	}
	fmt.Println("YES")
}
