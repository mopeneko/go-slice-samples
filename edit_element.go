package main

import "fmt"

func main() {
	s := []int{1, 2, 3}
	s[0] = 0
	fmt.Printf("%+v\n", s) // [0 2 3]
}