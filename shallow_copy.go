package main

import "fmt"

func main() {
	s1 := []int{1, 2, 3}
	s2 := s1
	s2[0] = 0
	fmt.Printf("s1=%+v s2=%+v\n", s1, s2) // s1=[0 2 3] s2=[0 2 3]
}