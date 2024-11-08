package main

import "fmt"

func main() {
	var n int
	fmt.Scan(&n)

	graph := make([][]int, n)
	for i := 0; i < n; i++ {
		graph[i] = make([]int, n)
		for j := 0; j < n; j++ {
			var v int
			fmt.Scan(&v)
			if i == j {
				continue
			}
			if v == 1 {
				graph[i][j] = 0
				continue
			}
			if v == 0 {
				graph[i][j] = 1
				continue
			}
		}
	}

	for i := 0; i < n; i++ {
		for j := 0; j < n; j++ {
			fmt.Printf("%d ", graph[i][j])
		}
		fmt.Print("\n")
	}
}
