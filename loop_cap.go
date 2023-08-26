package main

import "fmt"

func main() {
	s := make([]int, 0, 3)
	for i := 0; i < 3; i++ {
		if i == 1 {
			continue
		}
		s = append(s, i)
	}
	PrintSliceInfo(s) // [0 2] | len=2 cap=3
}

func PrintSliceInfo[T any](s []T) {
	fmt.Printf("%+v | len=%d cap=%d", s, len(s), cap(s))
}