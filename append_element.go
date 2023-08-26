package main

import "fmt"

func main() {
	s1 := make([]int, 0, 1)
	PrintSliceInfo(s1) // [] | len=0 cap=1 ptr=0xc0000a8000
	s1 = append(s1, 1)
	PrintSliceInfo(s1) // [1] | len=1 cap=1 ptr=0xc0000a8000

	s2 := make([]int, 0)
	PrintSliceInfo(s1) // [1] | len=1 cap=1 ptr=0xc0000a8000
	s2 = append(s2, 1)
	PrintSliceInfo(s2) // [1] | len=1 cap=1 ptr=0xc0000a8030
}

func PrintSliceInfo[T any](s []T) {
	fmt.Printf("%+v | len=%d cap=%d ptr=%p\n", s, len(s), cap(s), s)
}