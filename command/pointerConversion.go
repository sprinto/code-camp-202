package main

import (
	"fmt"
	"unsafe"
)

type Test struct {
	str string
	num int64
}

func main() {
	t := Test{str: "hello", num: 1}
	tp := unsafe.Pointer(&t)
	tsp := (*string)(unsafe.Pointer(tp))
	*tsp = "world"
	tip := (*int64)(unsafe.Pointer(uintptr(tp) + unsafe.Offsetof(t.num)))
	*tip = 2
	fmt.Printf("t.str: %s, t.num: %d", t.str, t.num)
}
