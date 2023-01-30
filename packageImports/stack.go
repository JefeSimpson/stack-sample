package pkgimpt

import (
	"fmt"
	"os"
)

type Node struct {
	data int
	prev *Node
}

type Stack struct {
	top  *Node
	size int
}

func (s *Stack) Pop() int {
	if s.size == 0 {
		os.Exit(3)
		//return 0
	}
	dt := s.top.data
	s.top = s.top.prev
	s.size--
	return dt
}

func (s *Stack) Clear() {
	i := 0
	for i < s.size {
		s.Pop()
	}
}

func (s *Stack) Insert(data int) {
	node := &Node{data, s.top}
	s.top = node
	s.size++
}

func (s *Stack) Top() int {
	if s.size == 0 {
		os.Exit(3)
	}
	return s.top.data
}

func (s *Stack) Contains(data int) bool {
	newNode := s.top
	for newNode != nil {
		if newNode.data == data {
			return true
		}
		newNode = newNode.prev
	}
	return false
}

func (s *Stack) Increment() {
	newNode := s.top
	for newNode != nil {
		newNode.data++
		newNode = newNode.prev
	}
}

func (s *Stack) Print() {
	newNode := s.top
	for newNode != nil {
		fmt.Print(newNode.data, " ")
		newNode = newNode.prev
	}
}

func (s *Stack) PrintReverse() {
	st := &Stack{}
	newNode := s.top
	for newNode != nil {
		st.Insert(newNode.data)
		newNode = newNode.prev
	}
	for st.size != 0 {
		fmt.Print(st.Pop(), " ")
	}
	fmt.Println()
}

func (s *Stack) Size() int {
	return s.size
}
