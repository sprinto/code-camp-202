package main

import (
	"fmt"
)

type TheCb func(i1 int) int

type NcSession struct {
	User        string
	Pwd         string
	ncSessionId int
	NcCallback  TheCb
}

func main() {
	var s = NcSession{"tomas", "hemligt", 0, counter}

	i := s.NcCallback(3)
	fmt.Printf("cb ==> %d\n", i)
}

func counter(i int) int {
	fmt.Println("=== Callback working ==")
	return i * i
}
