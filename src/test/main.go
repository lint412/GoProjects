// test project main.go
package main

import (
	//	"errors"
	//	"bytes"
	"fmt"
	//	"math"
	//	"runtime"
	//	"unsafe"
	//"os"
	//"strconv"
	//	"io"
	//	"log"
)

//var num int = 10
//var numx2, numx3 int

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

	/*for i, j, s := 0, 5, "a"; i < 3 && j < 100 && s != "aaaaa"; i, j,
		s = i+1, j+1, s+"a" {
		fmt.Println("Value of i, j, s:", i, j, s)
	}*/

	//	numx2, numx3 = getX2AndX3(num)
	//	PrintValues()
	//	numx2, numx3 = getX2AndX3_2(num)
	//	PrintValues()

	//	sum, calc, minus := operateVal(20, 9)
	//	fmt.Printf("return value are: %d %d %d", sum, calc, minus)

	//	fmt.Print("First example with -1: ")
	//	ret1, err1 := MySqrt(-1)
	//	if err1 != nil {
	//		fmt.Println("Error! Return values are: ", ret1, err1)
	//	} else {
	//		fmt.Println("It's ok! Return values are: ", ret1, err1)
	//	}

	//	fmt.Print("Second example with 5: ")
	//	//you could also write it like this
	//	if ret2, err2 := MySqrt(5); err2 != nil {
	//		fmt.Println("Error! Return values are: ", ret2, err2)
	//	} else {
	//		fmt.Println("It's ok! Return values are: ", ret2, err2)
	//	}
	//	// named return variables:
	//	fmt.Println(MySqrt2(5))

	//	n := 0
	//	reply := &n
	//	Multiply(10, 5, reply)
	//	fmt.Println("Multiply:", n) // Multiply: 50

	//	x := min(1, 3, 2, 0)
	//	fmt.Printf("The minimum is: %d\n", x)
	//	arr := []int{7, 9, 3, 5, 1}
	//	x = min(arr...)
	//	fmt.Printf("The minimum in the array arr is: %d", x)

	//	a()

	//	func1("Go")

	//	printValue(0)

	//	func(cn int) {
	//		sum := 0
	//		for i := 1; i <= cn; i++ {
	//			sum += i
	//		}
	//		fmt.Print(sum)
	//	}(5)

	//	f := adder(5)
	//	fmt.Println("result=", f(3))
	//	fmt.Println("result=", f(2))

	//	var arrKeyValue = []string{1: "Chris", 3: "Ron"}
	//	fmt.Println(len(arrKeyValue))

	//	s := make([]byte, 5)
	//	fmt.Println(len(s), " ", cap(s))
	//	s = s[2:4]
	//	fmt.Println(len(s), " ", cap(s))
	//	s1 := []byte{'p', 'o', 'e', 'm'}
	//	s2 := s1[2:]
	//	fmt.Println(len(s2), " ", cap(s2))
	//	for i := 0; i < len(s2); i++ {
	//		fmt.Printf("%c", s2[i])
	//	}
	//	fmt.Println()
	//	s2[1] = 't'

	//	for i := 0; i < len(s1); i++ {
	//		fmt.Printf("%c\n", s1[i])
	//	}
	//	fmt.Println()

	//	s3 := bytes.SplitN(s1, s1[2:3], 2)
	//	//s3 := bytes.Split(s1, nil)
	//	//bytes.Buffer.ReadRune()

	//	for i := 0; i < len(s3); i++ {
	//		for j := 0; j < len(s3[i]); j++ {
	//			fmt.Printf("%c\n", s3[i][j])
	//		}
	//		fmt.Println()
	//	}

	//	items := [...]int{10, 20, 30, 40, 50}
	//	for i, _ := range items {
	//		items[i] *= 2
	//	}

	//	for _, item := range items {
	//		fmt.Println(item)
	//	}

	//	sl := items[5:6]
	//	fmt.Println(len(sl), " ", cap(sl))

	sl_from := []int{1, 2, 3}
	sl_to := make([]int, 10)

	n := copy(sl_to, sl_from)
	fmt.Println(sl_to)
	fmt.Printf("Copied %d elements\n", n) // n == 3

	sl3 := []int{1, 2, 3}
	sl3 = append(sl3, 4, 5, 6)
	fmt.Println(sl3)

}

//func PrintValues() {
//	fmt.Printf("num = %d, 2x num = %d, 3x num = %d\n", num, numx2, numx3)
//}

//func getX2AndX3(input int) (int, int) {
//	return 2 * input, 3 * input
//}

//func getX2AndX3_2(input int) (x2 int, x3 int) {
//	x2 = 2 * input
//	x3 = 3 * input
//	//	return x2, x3
//	return
//}

//func operateVal(a, b int) (int, int, int) {
//	return a + b, a * b, a - b
//}

//func operateVal(a, b int) (sum, cal, minu int) {
//	sum, cal, minu = a+b, a*b, a-b
//	return
//}

//func MySqrt(f float64) (float64, error) {
//	//return an error as second parameter if invalid input
//	if f < 0 {
//		return float64(math.NaN()), errors.New("I won't be able to do a sqrt of negative number!")
//	}
//	//otherwise use default square root function
//	return math.Sqrt(f), nil
//}

////name the return variables - by default it will have 'zero-ed' values i.e. numbers are 0, string is empty, etc.
//func MySqrt2(f float64) (ret float64, err error) {
//	if f < 0 {
//		//then you can use those variables in code
//		ret = float64(math.NaN())
//		err = errors.New("I won't be able to do a sqrt of negative number!")
//	} else {
//		ret = math.Sqrt(f)
//		//err is not assigned, so it gets default value nil
//	}
//	//automatically return the named return variables ret and err
//	return
//}

//func Multiply(a, b int, reply *int) {
//	*reply = a * b
//}

//func min(a ...int) int {
//	if len(a) == 0 {
//		return 0
//	}
//	min := a[0]
//	for _, v := range a {
//		fmt.Println(v)
//		if v < min {
//			min = v
//		}
//	}
//	return min
//}

//func a() {
//	i := 1
//	defer fmt.Println(i)
//	defer fmt.Println(9)
//	i++
//	fmt.Printf("last value: %d\n", i)
//	return
//}

//func enter(s string, n int, err error) {
//	log.Printf("func1(%q) = %d, %v", s, n, err)
//}

//func func1(s string) (n int, err error) {
//	defer func() {
//		log.Printf("func1(%q) = %d, %v", s, n, err)
//	}()
//	//	defer enter(s, n, err)
//	s = "Other"
//	return 7, io.EOF
//}

//func printValue(i int) {
//	i++
//	if i > 10 {
//		return
//	}
//	printValue(i)
//	fmt.Println(i)
//}

//func adder(a int) func(b int) int {
//	fmt.Println("a=", a)
//	return func(b int) int {
//		a++
//		return a + b
//	}
//}
