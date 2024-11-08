package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			var e int
			fmt.Scan(&e)
			graph[i][j] = e
		}
	}
	for i := 0; i < n; i++ {
		for j := i + 1; j < n; j++ {
			if graph[i][j] == 1 {
				fmt.Printf("%d %d\n", i+1, j+1)
			}
		}
	}
}
