package main

import (
	"fmt"

	"./other"
)

func init() {
	fmt.Println("world init: will it exec?")
}
func otherFileInSamePackage() {
	fmt.Println("This is Hello World" + " from world in same package main")
	fmt.Printf("other Boat: %s\n", other.Boat)

	fmt.Printf("var from hello - home: %s, user: %s\n", HOME, USER)
}
