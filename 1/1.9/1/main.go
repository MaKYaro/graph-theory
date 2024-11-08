package main

import (
	"fmt"
)

func main() {
	s := Stack{[]int{}}
	for {
		var command string
		var number int
		fmt.Scanf("%s %d", &command, &number)

		switch command {
		case "push":
			s.Push(number)
			fmt.Println("ok")
		case "pop":
			val := s.Pop()
			fmt.Println(val)
		case "back":
			val := s.Back()
			fmt.Println(val)
		case "size":
			val := s.Size()
			fmt.Println(val)
		case "clear":
			s.Clear()
			fmt.Println("ok")
		case "exit":
			fmt.Println("bye")
			return
		}

	}
}

type Stack struct {
	items []int
}

func (s *Stack) Push(n int) {
	s.items = append(s.items, n)
}

func (s *Stack) Pop() int {
	lastElem := s.items[len(s.items)-1]
	s.items = s.items[:len(s.items)-1]

	return lastElem
}

func (s *Stack) Back() int {
	return s.items[len(s.items)-1]
}

func (s *Stack) Size() int {
	return len(s.items)
}

func (s *Stack) Clear() {
	s.items = s.items[:0]
}
