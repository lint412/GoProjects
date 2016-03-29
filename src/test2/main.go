// test2 project main.go
package main

import (
	//	"encoding/gob"
	//	"crypto/sha1"
	//"bytes"
	"fmt"
	"time"
	//	"io"
	//	"log"
	//	"os"
	//	"strings"
	//	"flag"
	//	"math"
	//	"runtime/pprof"
	//	"runtime"
)

//var cpuprofile = flag.String("cpuprofile", "", "write cpu profile to file")
//var numCores = flag.Int("n", 2, "number of CPU cores to use")

//var (
//	firstName, lastName, s string
//	i                      int
//	f                      float32
//	input                  = "56.12 / 5212 / Go"
//	format                 = "%f / %d / %s"
//)

//var inputReader *bufio.Reader
//var input string
//var err error

//type Page struct {
//	Title string
//	Body  []byte
//}

//func (self *Page) save() {
//	outputFile, outputError := os.OpenFile(self.Title+".txt", os.O_WRONLY|os.O_CREATE, 0666)
//	if outputError != nil {
//		fmt.Printf("An error occurred with file opening or creation\n")
//		return
//	}
//	defer outputFile.Close()

//	outputWriter := bufio.NewWriter(outputFile)
//	outputWriter.Write(self.Body)
//	outputWriter.Flush()
//}

//func (self *Page) load(title string) {
//	inputFile, inputError := os.OpenFile(title, os.O_RDONLY, 0)
//	if inputError != nil {
//		fmt.Printf("An error occurred with file opening\n")
//		return
//	}
//	defer inputFile.Close()

//	self.Title = title

//	fmt.Println(self.Title)

//	inputReader := bufio.NewReader(inputFile)
//	for {

//		inputBytes, _, readerError := inputReader.ReadLine()
//		//inputString, readerError := inputReader.ReadString('\n')
//		//fmt.Println(inputBytes)
//		if readerError == io.EOF {
//			break
//		}
//		inputString := string(inputBytes) + "\r\n"
//		fmt.Println(inputString)
//		self.Body = append(self.Body, inputString...)
//	}

//	fmt.Printf("after load: %s", self.Body)
//}

//type Address struct {
//	Type    string
//	City    string
//	Country string
//}

//type VCard struct {
//	FirstName string
//	LastName  string
//	Addresses []*Address
//	Remark    string
//}

//func longWait() {
//	fmt.Println("Beginning longWait()")
//	time.Sleep(5 * 1e9) // sleep for 5 seconds
//	fmt.Println("End of longWait()")
//}

//func shortWait() {
//	fmt.Println("Beginning shortWait()")
//	time.Sleep(2 * 1e9) // sleep for 2 seconds
//	fmt.Println("End of shortWait()")
//}

func sendData(ch chan string) {
	ch <- "Washington"
	//	fmt.Println("1")
	ch <- "Tripoli"
	//	fmt.Println("2")
	ch <- "London"
	//	fmt.Println("3")
	ch <- "Beijing"
	ch <- "Tokio"
	fmt.Println("before close")
	close(ch)
	fmt.Println("after close")
}

func getData(ch chan string) {
	//	var input string
	//	// time.Sleep(1e9)
	//	input = <-ch
	//	fmt.Printf("%s ", input)
	//	input = <-ch
	//	fmt.Printf("%s ", input)
	//	for {
	//		input = <-ch
	//		fmt.Printf("%s ", input)
	//	}

	time.Sleep(1e9)

	//	for {
	//		input, open := <-ch
	//		if !open {
	//			break
	//		}
	//		fmt.Printf("%s ", input)
	//	}
	for input := range ch {
		fmt.Printf("%s ", input)
	}
}

//func f1(in chan int) {
//	fmt.Println(<-in)
//}

func main() {
	//	fmt.Println("Please enter your full name: ")
	//	fmt.Scanln(&firstName, &lastName)
	//	// fmt.Scanf("%s %s", &firstName, &lastName)
	//	fmt.Printf("Hi %s %s!\n", firstName, lastName) // Hi Chris Naegels
	//	fmt.Sscanf(input, format, &f, &i, &s)
	//	fmt.Println("From the string we read: ", f, i, s)
	//	// 输出结果: From the string we read: 56.12 5212 Go

	//	inputReader = bufio.NewReader(os.Stdin)
	//	fmt.Println("Please enter some input: ")
	//	input, err = inputReader.ReadString('\n')
	//	if err == nil {
	//		fmt.Printf("The input was: %s\n", input)
	//	}

	//	inputReader := bufio.NewReader(os.Stdin)
	//	fmt.Println("Please enter your name:")
	//	input, err := inputReader.ReadString('\n')

	//	if err != nil {
	//		fmt.Println("There were errors reading, exiting program.")
	//		return
	//	}

	//	fmt.Printf("Your name is %s", input)
	//	// For Unix: test with delimiter "\n", for Windows: test with "\r\n"
	//	switch input {
	//	case "Philip\r\n":
	//		fmt.Println("Welcome Philip!")
	//	case "Chris\r\n":
	//		fmt.Println("Welcome Chris!")
	//	case "Ivo\r\n":
	//		fmt.Println("Welcome Ivo!")
	//	default:
	//		fmt.Printf("You are not welcome here! Goodbye!")
	//	}

	//	inputReader := bufio.NewReader(os.Stdin)
	//	for {
	//		fmt.Println("Please enter letters:")
	//		input, err := inputReader.ReadString('S')

	//		if err != nil {
	//			fmt.Println("There were errors reading, exiting program.")
	//			return
	//		}

	//		fmt.Printf("Your letters are %s\n", input)

	//		fmt.Println(len(input))

	//		s := strings.Replace(input, "\r\n", "", -1)

	//		sword := strings.Replace(input, "\r\n", " ", -1)
	//		words := strings.Split(sword, " ")

	//		fmt.Println("letters count: ", len(s))

	//		fmt.Println("words count: ", len(words))
	//	}

	//	ap := new(Page)
	//	ap.Title = "test file"
	//	s := "hello world!\r\nhello word2"
	//	c := []byte(s)
	//	ap.Body = c

	//	fmt.Printf("%s", ap.Body)

	//	ap.save()

	//	ap.load("test file.txt")

	//	argstring := strings.Join(os.Args[:], " ")
	//	fmt.Println(argstring)

	//	inputFile, _ := os.Open("main.go")
	//	outputFile, _ := os.OpenFile("goprogramT.go", os.O_WRONLY|os.O_CREATE, 0666)
	//	defer inputFile.Close()
	//	defer outputFile.Close()
	//	inputReader := bufio.NewReader(inputFile)
	//	outputWriter := bufio.NewWriter(outputFile)
	//	for {
	//		inputString, _, readerError := inputReader.ReadLine()
	//		if readerError == io.EOF {

	//			fmt.Println("EOF")
	//			break
	//		}
	//		outputString := string([]byte(inputString)[2:5]) + "\r\n"
	//		_, err := outputWriter.WriteString(outputString)
	//		if err != nil {
	//			fmt.Println(err)
	//			break
	//		}
	//	}
	//	outputWriter.Flush()
	//	fmt.Println("Conversion done")

	//	pa := &Address{"private", "Aartselaar", "Belgium"}
	//	wa := &Address{"work", "Boom", "Belgium"}
	//	vc := VCard{"Jan", "Kersschot", []*Address{pa, wa}, "none"}
	//	fmt.Printf("%v: \n", vc) // {Jan Kersschot [0x126d2b80 0x126d2be0] none}:
	//	fmt.Printf("%s: \n", vc)
	// JSON format:
	//	js, _ := json.Marshal(vc)
	//	fmt.Printf("JSON format: %s", js)
	//	// using an encoder:
	//	file, _ := os.OpenFile("vcard.json", os.O_CREATE|os.O_WRONLY, 0)
	//	defer file.Close()
	//	enc := json.NewEncoder(file)
	//	err := enc.Encode(vc)
	//	if err != nil {
	//		log.Println("Error in encoding json")
	//	}

	//	file, _ := os.OpenFile("vcard.gob", os.O_CREATE|os.O_WRONLY, 0666)
	//	defer file.Close()
	//	enc := gob.NewEncoder(file)
	//	err := enc.Encode(vc)
	//	if err != nil {
	//		log.Println("Error in encoding gob")
	//	}

	//	flag.Parse()
	//	if *cpuprofile != "" {
	//		f, err := os.Create(*cpuprofile)
	//		if err != nil {
	//			log.Fatal(err)
	//		}
	//		pprof.StartCPUProfile(f)
	//		defer pprof.StopCPUProfile()
	//	}

	//	hasher := sha1.New()
	//	for i := 0; i < math.MaxUint16; i++ {
	//		io.WriteString(hasher, "test")
	//		b := []byte{}
	//		fmt.Printf("Result: %x\n", hasher.Sum(b))
	//		fmt.Printf("Result: %d\n", hasher.Sum(b))

	//		hasher.Reset()
	//		data := []byte("We shall overcome!")
	//		n, err := hasher.Write(data)
	//		if n != len(data) || err != nil {
	//			log.Printf("Hash write error: %v / %v", n, err)
	//		}
	//		checksum := hasher.Sum(b)
	//		fmt.Sprintf("Result: %x\n", checksum)

	//	}

	//	flag.Parse()
	//	fmt.Println("core: ", *numCores)
	//	runtime.GOMAXPROCS(*numCores)

	//	start := time.Now()
	//	fmt.Println("In main()")
	//	go longWait()
	//	go longWait()
	//	go longWait()
	//	go shortWait()

	//	fmt.Println("About to sleep in main()")
	//	// sleep works with a Duration in nanoseconds (ns) !
	//	time.Sleep(10 * 1e9)
	//	fmt.Println("At the end of main()")
	//	end := time.Now()
	//	delta := end.Sub(start)
	//	fmt.Println("Time: ", delta)

	//	ch := make(chan string)

	//	go sendData(ch)
	//	go getData(ch)

	//	time.Sleep(3 * 1e9)
	//	fmt.Println("Finish")

	//	out := make(chan int)
	//	//	out <- 5
	//	go f1(out)
	//	out <- 5

	ch := make(chan string)

	go sendData(ch)
	getData(ch)

}
