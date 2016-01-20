// fibutil project main.go
package main

import (
	//	"fibutil/fib"
	"fmt"
	"os"
	//	"strconv"
)

func main() {

	//	for i := 0; i < 10; i++ {
	//		fmt.Println(fib.Fibm(int64(i)))
	//	}

	goosInfo()

	//	switch len(os.Args) {
	//	case 1:
	//		fmt.Println(os.Args[0])
	//	case 2:
	//		n, err := strconv.ParseInt(os.Args[1], 10, 64)
	//		if err == nil {
	//			fmt.Println(fib.Fibm(n))
	//			return
	//		}
	//	case 3:
	//		n1, e1 := strconv.ParseInt(os.Args[1], 10, 64)
	//		n2, e2 := strconv.ParseInt(os.Args[2], 10, 64)
	//		if e1 == nil && e2 == nil {
	//			fmt.Println(fib.Fibm(n1))
	//			fmt.Println(fib.Fibm(n2))
	//			return
	//		}
	//	}
	//	fmt.Fprintf(os.Stderr, "%s, fibonacci number util\n\tfibutil n\t:fibonacci number\n\tfibutil n1 n2\t:fibonacci number list\n", os.Args[0])
}

func goosInfo() {
	var goos string = os.Getenv("GOOS")
	fmt.Printf("The operating system is: %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is: %s\n", path)
}
