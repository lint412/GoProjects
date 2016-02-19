// test project main.go
package main

import (
	"fmt"
	//	"runtime"
	//	"unsafe"
	//"os"
	//"strconv"
)

/*var prompt = "Enter a digit, e.g. 3 " + "or %s to quit."

func init() {
	if runtime.GOOS == "windows" {
		prompt = fmt.Sprintf(prompt, "Ctrl+Z, Enter")
	} else { //Unix-like
		prompt = fmt.Sprintf(prompt, "Ctrl+D")
	}
	fmt.Println(prompt)
}*/

func main() {
	/*
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

		if i1 == 5 {
			i1 = 6
		} else if i1 > 5 {
			i1 = 7
		}
	*/

	/*var orig string = "ABC"
	// var an int
	var newS string
	// var err error
	fmt.Printf("The size of ints is: %d\n", strconv.IntSize)
	// anInt, err = strconv.Atoi(origStr)
	an, err := strconv.Atoi(orig)
	if err != nil {
		//fmt.Printf("orig %s is not an integer - exiting with error\n", orig)
		//return
		fmt.Printf("Program stopping with error %v", err)
		os.Exit(1)
	}
	fmt.Printf("The integer is %d\n", an)
	an = an + 5
	newS = strconv.Itoa(an)
	fmt.Printf("The new string is: %s\n", newS)*/

	/*k := 6
	switch k {
	case 4:
		fmt.Println("was <= 4")
		fallthrough
	case 5:
		fmt.Println("was <= 5")
		fallthrough
	case 6:
		fmt.Println("was <= 6")
		fallthrough
	case 7:
		fmt.Println("was <= 7")
		fallthrough
	case 8:
		fmt.Println("was <= 8")
		fallthrough
	default:
		fmt.Println("default case")
	}*/

	/*var season string
	mon := 19
	switch mon {
	case 1, 2, 3:
		season = "春天"
	case 4, 5, 6:
		season = "夏天"
	case 7, 8, 9:
		season = "秋天"
	case 10, 11, 12:
		season = "冬天"
	default:
		fmt.Println("无效月份")
		return
	}
	fmt.Println(season)*/

	/*for i := 0; i < 15; i++ {
		fmt.Println(i)
	}*/

	/*idx := 0
	HERE:
		{
			fmt.Println(idx)
			idx++
			if idx < 15 {
				goto HERE
			}
		}*/

	/*for i := 0; i < 25; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("G")
		}
		fmt.Print("\n")
	}

	var strG string = ""
	for i := 0; i < 25; i++ {
		strG = strG + "G"
		fmt.Println(strG)
	}*/

	/*for i := 0; i <= 10; i++ {
		fmt.Printf("pos %b is %b\n", i, ^i)
	}*/

	/*for i := 1; i <= 100; i++ {
		switch {
		case i%3 == 0 && i%5 == 0:
			fmt.Println("FizzBuzz")
		case i%3 == 0:
			fmt.Println("Fizz")
		case i%5 == 0:
			fmt.Println("Buzz")
		default:
			fmt.Println(i)
		}
	}*/

	/*for i := 0; i < 10; i++ {
		for j := 0; j < 20; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}*/

	/*str := "Go is a beautiful language!"
	fmt.Printf("The length of str is: %d\n", len(str))
	for pos, char := range str {
		fmt.Printf("Character on position %d is: %c \n", pos, char)
	}
	fmt.Println()
	str2 := "Chinese: 日本語"
	fmt.Printf("The length of str2 is: %d\n", len(str2))
	for pos, char := range str2 {
		//fmt.Printf("character %c starts at byte position %d\n", char, pos)
		fmt.Printf("%-2d      %d      %U '%c' % X\n", pos, char, char, char, []byte(string(char)))
	}
	fmt.Println()
	fmt.Println("index int(rune) rune    char bytes")
	for index, rune := range str2 {
		fmt.Printf("%-2d      %d      %U '%c' % X\n", index, rune, rune, rune, []byte(string(rune)))
	}*/

	/*for i := 0; i < 5; i++ {
		var v int
		fmt.Printf("%d ", v)
		v = 5
	}*/

	for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}

}
