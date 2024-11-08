package main

import "fmt"

// func main() {
// 	var n, k int
// 	fmt.Scan(&n, &k)

// 	if n < k {
// 		fmt.Println(n)
// 	} else {
// 		fmt.Println(k / 2 + 1)
// 	}
// }

const (
	UNVISITED = 0
	VISITING  = 1
	VISITED   = 2
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	result := dfs(n, k)

	fmt.Println(result)
}

func dfs(n, k int) int {
	if n == k {
		return n/2 + 1
	}
	states := make([]int, n+1)
	result := 0

	for v := 1; v <= n; v++ {
		state := states[v]
		if state == UNVISITED {
			result++
			dfsVisit(states, v, n, k)
		}
	}
	return result
}

func dfsVisit(states []int, v, n, k int) {
	states[v] = VISITED
	step := 1

	for ; step < k; step++ {
		if (2*v+step)%k == 0 {
			break
		}
	}
	v = v + step
	for v <= n {
		states[v] = VISITED
		step = k - step
		v = v + step
	}
}
