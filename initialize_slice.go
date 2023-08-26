package main

import "fmt"

func main() {
	s1 := []int{}
	PrintSliceInfo(s1) // [] | len=0 cap=0 isNil=false

	s2 := []int{1, 2, 3}
	PrintSliceInfo(s2) // [1 2 3] | len=3 cap=3 isNil=false

	s3 := make([]int, 0)
	PrintSliceInfo(s3) // [] | len=0 cap=0 isNil=false
}

func PrintSliceInfo[T any](s []T) {
	fmt.Printf("%+v | len=%d cap=%d isNil=%t\n", s, len(s), cap(s), s == nil)
}