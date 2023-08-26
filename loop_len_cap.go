package main

import "fmt"

func main() {
	s := make([]int, 3)
	for i := 0; i < 3; i++ {
		s[i] = i
	}
	PrintSliceInfo(s) // [0 1 2] | len=3 cap=3
}

func PrintSliceInfo[T any](s []T) {
	fmt.Printf("%+v | len=%d cap=%d", s, len(s), cap(s))
}