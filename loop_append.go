package main

import "fmt"

func main() {
	s := make([]int, 0)
	for i := 0; i < 3; i++ {
		s = append(s, i)
	}
	PrintSliceInfo(s) // [0 1 2] | len=3 cap=4
}

func PrintSliceInfo[T any](s []T) {
	fmt.Printf("%+v | len=%d cap=%d", s, len(s), cap(s))
}