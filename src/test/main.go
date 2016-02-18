// test project main.go
package main

import (
	"fmt"
	"unsafe"
)

func main() {
	//	fmt.Printf("Sizeof i1 is %d byte(s)\n", unsafe.Sizeof(i1))
	var i1 = 5
	fmt.Printf("An integer: %d, its location in memory: %p\n", i1, &i1)
	var intP *int
	intP = &i1
	fmt.Printf("The value at memory location %p is %d\n", intP, *intP)
	fmt.Printf("Sizeof intP is %d byte(s)\n", unsafe.Sizeof(intP))

	s := "good bye"
	var p *string = &s
	*p = "ciao"
	fmt.Printf("Here is the pointer p: %p\n", p)  // prints address
	fmt.Printf("Here is the string *p: %s\n", *p) // prints string
	fmt.Printf("Here is the string s: %s\n", s)   // prints same string
	fmt.Printf("Sizeof p is %d byte(s)\n", unsafe.Sizeof(p))
}
