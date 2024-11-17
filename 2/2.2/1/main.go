package main

import "fmt"

func main() {
	var n, m int
	fmt.Scan(&n, &m)

	cityOwner := make([]int, n+1)
	for i := 1; i <= n; i++ {
		var owner int
		fmt.Scan(&owner)
		cityOwner[i] = owner
	}

	roadGraph := make(map[int][]Road, n)
	for i := 1; i <= n; i++ {
		roadGraph[i] = make([]Road, 0)
	}
	for i := 0; i < m; i++ {
		var c1, c2 int
		fmt.Scan(&c1, &c2)

		if cityOwner[c1] == cityOwner[c2] {
			road1 := NewRoad(c2, 0)
			road2 := NewRoad(c1, 0)
			roadGraph[c1] = append(roadGraph[c1], road1)
			roadGraph[c2] = append(roadGraph[c2], road2)
		} else {
			road1 := NewRoad(c2, 1)
			road2 := NewRoad(c1, 1)
			roadGraph[c1] = append(roadGraph[c1], road1)
			roadGraph[c2] = append(roadGraph[c2], road2)
		}
	}

	price, roadLen := bfs(roadGraph, n)
	if price == -1 {
		fmt.Println("impossible")
	} else {
		fmt.Println(price, roadLen)
	}
}

const (
	UNVISITED = 0
	VISITING  = 1
	VISITED   = 2

	FREE_ROAD = 0
	TOLL_ROAD = 1
)

type Road struct {
	City, Cost int
}

func NewRoad(city, cost int) Road {
	return Road{city, cost}
}

type Deque struct {
	items []int
}

func NewDeque(size int) *Deque {
	items := make([]int, 0, size)

	return &Deque{items: items}
}

func (d *Deque) PushFront(item int) {
	d.items = append([]int{item}, d.items...)
}

func (d *Deque) PushBack(item int) {
	d.items = append(d.items, item)
}

func (d *Deque) PopFront() int {
	frontElement := d.items[0]
	d.items = d.items[1:]

	return frontElement
}

func (d *Deque) PopBack() int {
	backElement := d.items[d.Size()-1]
	d.items = d.items[:d.Size()-1]

	return backElement
}

func (d *Deque) Size() int {
	return len(d.items)
}

func bfs(roadGraph map[int][]Road, n int) (int, int) {
	states := make([]int, n+1)
	states[1] = VISITING

	d := NewDeque(n)
	d.PushBack(1)

	price := 0
	visitedCityNum := 0
	for d.Size() > 0 {
		l := d.Size()

		for l > 0 {
			curr := d.PopFront()
			visitedCityNum++
			for _, nei := range findUnvisitedNei(roadGraph, states, n, curr) {
				//fmt.Println(nei)
				if nei.City == n {
					visitedCityNum++
					return price + nei.Cost, visitedCityNum
				}
				if nei.Cost == FREE_ROAD {
					d.PushFront(nei.City)
					l++
				}
				if nei.Cost == TOLL_ROAD {
					d.PushBack(nei.City)
				}
				states[nei.City] = VISITING
			}
			states[curr] = VISITED
			l--
		}
		price++
	}

	return -1, -1
}

func findUnvisitedNei(graph map[int][]Road, states []int, n, v int) []Road {
	result := make([]Road, 0, n)
	for _, nei := range graph[v] {
		if states[nei.City] == UNVISITED {
			result = append(result, nei)
		}
	}

	return result
}
