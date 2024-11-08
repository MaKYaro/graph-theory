package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	edges := make([][2]int, m)
	for i := 0; i < m; i++ {
		var v1, v2 int
		fmt.Scan(&v1, &v2)
		edges[i] = [2]int{v1, v2}
	}

	var k int
	fmt.Scan(&k)

	edgesList := makeEdgesList(edges, n)
	result := countNeighboringEdges(edgesList, edges[k-1])

	fmt.Println(result)
}

func countNeighboringEdges(
	edgesList map[int][]int,
	edge [2]int,
) int {
	v1, v2 := edge[0], edge[1]

	v1NeighborsAmount := len(edgesList[v1]) - 1
	v2NeighborsAmount := len(edgesList[v2]) - 1

	return v1NeighborsAmount + v2NeighborsAmount
}

func makeEdgesList(edges [][2]int, n int) map[int][]int {
	edgesList := make(map[int][]int, n)
	for _, edge := range edges {
		v1, v2 := edge[0], edge[1]
		edgesList[v1] = append(edgesList[v1], v2)
		edgesList[v2] = append(edgesList[v2], v1)
	}

	return edgesList
}
