package main

import (
	"fmt"
	"os"

	"rsc.io/quote"
)

const c = "C"

// enumeration without type name
const (
	Unknown = 0
	Female  = 1
	Male    = 2
)

var v int = 5
var (
	HOME = os.Getenv("HOME")
	USER = os.Getenv("USER")
)

type T struct{}

// enumeration with type name
type Nummer int

const (
	ONE Nummer = iota
	TWO
	THREE
)

func init() { // initialization of package
	fmt.Printf("hello init: Now this package is initialized - a = %d, c = %s\n", v, c)
	v = v + 3
}

func main() {
	fmt.Println(quote.Go())
	fmt.Printf("After init - a = %d, c = %s\n", v, c)
	//otherFileInSamePackage()
	myOtherTypes()
}

func myOtherTypes() {
	var car []string
	fmt.Println("empty slice :", car)
}
func myTypes() {
	coral := [3]string{"blue coral", "staghorn coral", "pillar coral"}
	fmt.Println("fixed array ", coral)
	// coral = append(coral, "gul coral")   not w array

	seaCreatures := []string{"shark", "cuttlefish", "squid", "mantis shrimp"}
	fmt.Println("dynamic slice ", seaCreatures)
	seaCreatures = append(seaCreatures, "seahorse")
	fmt.Println("dynamic slice ", seaCreatures)

	sammy := map[string]string{"name": "Sammy", "animal": "shark", "color": "blue"}
	fmt.Println(sammy)
	fmt.Println(sammy["animal"])
}

func scanInput() {
	var name string
	fmt.Println("Ordet?")
	fmt.Scanln(&name)
	//	strings.TrimSpace(name)
	fmt.Printf("Ordet är ...<%s>\n", name)
}
