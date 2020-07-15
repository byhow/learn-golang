package main

import (
	"fmt"
	"math"
	"strconv"
	"reflect"
	"io/ioutil"
	"log"
	"net/http"
)
// when at this level, cannot infer type
var (
	j int = 99
	name string = "test"
)

// 1. if lowercase and scope to the package, same package can access
// 2. if uppercase and scope at this level, then exporting it makes
// 	  it globally visible
// 3. block scope

var (
	counter int = 0
)


const (
	a = iota
	b 
	c
)
// it will then enumerate automatically

const (
	a4 = iota
)

// iota usecase
const (
	errorSpecialist = iota
	// _ = iota
	// throw that away, like what it is in JS
	catSpecialist
	dogSpecialist
	snakeSpecialist
	// automatically enumerate on the fly
)

const (
	_ = iota
	KB = 1 << (10 * iota)
	MB
	GB
	TB
	PB
	EB
	ZB
	YB
)

const (
	isAdmin = 1 << iota
	isHeadquarters
	canSeeFinancials

	canSeeAfrica
	canSeeAsia
	canSeeEurope
	canSeeNorthAmerica
	canSeeSouthAmerica
)


type Doctor struct {
	number int
	actorName string
	companions []string
}

// composition or embedded struct
type Animal struct {
	Name string `required max: "100"`
	Origin string

}

type Bird struct {
	Animal
	Speed float32
	CanFly bool
}

func main() {
	// var i int = 42
	searchParameter := 28.5
	var theHTTPRequest string = "http"
	fmt.Printf(theHTTPRequest + "\n")
	var j int = int(searchParameter) // will lose information
	// variables can be declare only once under the SAME SCOPE
	// but if those are different scope, then its fine
	fmt.Printf("%v %T\n", j, j);
	
	var k int = 42
	var p string = strconv.Itoa(k) // will covert based on unicode
	fmt.Printf(p + "\n")


	// Primitives
	var n bool = false
	fmt.Printf("%v, %T\n", n, n)
	
	// m := 1 == 1
	var y bool = 1 == 2
	fmt.Printf("%v, %T\n", y, y)
	// var q bool // if just initialize, then every value has a zero value, which to bool is false

	// Numeric types
	// int will at least be 32 bits
	var w int16 = 42 // range to int64
	fmt.Printf("%v, %T\n", w, w)
	// var e uint16 = 42 // 0 - 65535, range to uint32

	// arithmetic operation
	fmt.Println(int(w) / k) // drop remainder
	// only allows same type, very strict on type conversion

	// Bitwise
	a := 10 // 1010
	b := 3 	// 0011
	fmt.Println(a & b) // 0010
	fmt.Println(a | b) // 1011
	fmt.Println(a ^ b) // 1001 only one is 1
	fmt.Println(a &^ b) // 0100 both 0

	// Bit shifting
	a = 8 // 2^3
	fmt.Println(a << 3) // 2^3 * 2^3 = 2^6
	fmt.Println(a >> 3) // 2^3 / 2^3 = 2^0

	// Floating point numbers
	na := 3.14
	na = 4.7e72
	na = 2.1E14
	fmt.Printf("%v, %T\n", na, na)

	// Complex number
	var nb complex64 = 1 + 2i
	fmt.Printf("%v, %T\n", real(nb), real(nb))
	fmt.Printf("%v, %T\n", imag(nb), imag(nb))

	var nb2 complex128 = complex(5, 12)
	fmt.Printf("%v, %T\n", nb2, nb2)

	// Text 
	s := "this is a string"
	// string is immutable
	fmt.Printf("%v, %T\n", s[6], s[6])
	// string concatenation is the same with other lang, +
	s1 := []byte(s)
	fmt.Printf("%v, %T\n", s1, s1) // got array of byte

	// rune type: UTF-32
	var r  rune = 'a' // use single quote
	fmt.Printf("%v, %T\n", r, r) // will get int32 type, just type alias for int32
	// we need this type for different apis, ReadRune vs ReadByte


	// Constants
	const myConst float64 = 1.57 // internal constant
	// would not allow constant to be calculated at compile time
	// const myConst float64 = math.Sin(1.57) // internal constant
	fmt.Printf("%v, %T\n", myConst, myConst)
	// has to be assignable at compile time

	const a1 = 42
	var b1 int16 = 27
	fmt.Printf("%v, %T\n", a1 + b1, a1 + b1) // this int + int16 works because the numbers add up as 42 + b ish

	// enumerated constant
	const a2 = iota

	// iota example storage
	fileSize := 4000000000.
	fmt.Printf("%2fGB\n", fileSize/GB)

	// iota admin bit operation (genius)
	// every flag is on their specific flag
	var roles byte = isAdmin | canSeeFinancials | canSeeEurope
	fmt.Printf("%b\n", roles)
	fmt.Printf("Is Admin? %v\n", isAdmin & roles == isAdmin) // bit mask 
	

	// arrays
	grades := [3]int{97, 98, 99} // [size]type{initializer}
	// grades := [...]int{9,10,11}
	var students [3]string
	students[0] = "Lisa"
	fmt.Printf("Students: %v\n", students)
	// continuous addrs
	fmt.Printf("Grades: %v\n", grades)
	// len(students) - got the size??

	// Identity matrix
	var im [3][3]int
	im[0] = [3]int {1, 0, 0}
	im[1] = [3]int {0, 1, 0}
	im[2] = [3]int {0, 0, 1}
	fmt.Println(im)

	// copying arrays are value copy
	a4 := [...]int{1,2,3}
	b2 := a4
	b2[2] = 5
	fmt.Println(a4)
	fmt.Println(b2)

	// if do pointers on addr, then its pointing 
	// e.g. b := &a

	a5 := []int{1, 2, 3} // slice: a projection to the underlying array
	fmt.Println(cap(a5)) // capacity function
	
	// literal slice
	// a6 := a5[:] // all elem
	// a6 = a5[1:] // from 2nd to end
	
	a6 := make([]int, 3)
	fmt.Println(a6)
	// get all zero values
	a6 = append(a6, 1, 2, 3, 4, 5, 6)
	fmt.Println(a6)
	// capacity will grow 

	// concatenation between slices has to be same type
	// spread operater ...
	a6 = append(a6, []int{2, 3, 4, 5}...)

	// stack operation
	// shift operation
	
	// Map 
	// the key has to be tested for equality
	// array is a valid key type but a slice is not
	ascii := map[string]int {
		"A": 65, 
		"B": 66,
	}

	// or use the make function
	// ascii := make(map[int]int)
	fmt.Println(ascii)

	ascii["C"] = 67
	delete(ascii, "C")
	// when requesting non-exist key, it will return 0, unless
	pop, ok := ascii["D"]
	fmt.Println(pop, ok) // ok is false if the key is not in the map
	// len(ascii) to find out the length 

	sp := ascii
	// manipulation affects other instance
	delete(sp, "A")
	fmt.Println(sp)
	fmt.Println(ascii)

	// Struct
	aDoctor := Doctor {
		number: 3,
		actorName: "John Doe",
		companions: []string {
			"Other",
			"People",
		},
	}
	fmt.Println(aDoctor.companions[1])
	// struct does not need to have the same data
	// a type of collection
	// always use this syntax instead of the positional syntax unless 
	// it is a short life-span struct

	// anonymous struct
	// those are independent datasets. copy will not affect the others
	// unless address operator is used
	bDoctor := struct{name string} {name: "John Doe"}
	fmt.Println(bDoctor)

	// composition instead of inheritance
	// just a syntactic sugar
	b4 := Bird{}
	b4.Name = "bird1"
	fmt.Println(b4)
	
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")
	fmt.Println(field.Tag) // string of tag inside struct


	// Control Flow
	if true {
		fmt.Println("The test is true")
	}

	// initializer block
	if pop, ok := ascii["E"]; ok {
		fmt.Println(pop)
	}

	// &&, ||, and or
	myNum := 0.123
	if math.Abs(myNum / math.Pow(math.Sqrt(myNum), 2) - 1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}


	i := 10
	switch {
	case i < 1:
		fmt.Println("This is 1")
	case i < 2:
		fmt.Println("This is 2")
		break // implied
	case i < 3:
		fmt.Println("This is 3")
		fallthrough // normal cases in C/C++
	case i < 10:
		fmt.Println("This is less than 10")
	default: 
		fmt.Println("This is default")
	}

	// type switch statement
	var i1 interface{} = 1.0
	switch i1.(type) {
	case int:
		fmt.Println("i is an int")
		break
		fmt.Println("needless")
	case float64:
		fmt.Println("i is float 64")
	case string:
		fmt.Println("i is a string")
	case [3]int:
		fmt.Println("i is a array of int of size 3")
	default:
		fmt.Println("i is an another type")
	}

	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	// no comma operators in Go
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		// increment operation i++ is not an expression
		fmt.Println(i, j)
	}

	// infinite loop

	ll := 0
	for {
		fmt.Println(1)
		ll++
		if ll==5 {
			break
		}
	}

	// for range loop
	s11 := []int{1, 2, 3}
	for k, v := range s11{
		fmt.Println(k, v)
	}
	
	// defer
	fmt.Println("Start")
	// when the function exits, it will look for defer
	// function that it left to call, and execute;
	// when it hits multiple defer functions, it executes 
	// in LIFO, last deferred, first to execute
	defer fmt.Println("Middle")
	fmt.Println("End")

	// common use case for defer
	// open and close 
	res, err := http.Get("http://google.com/robots.txt")
	if err != nil {
		log.Fatal(err)
	}
	// it is convenient, but when requesting a large chunk of body
	// not closing the resource may harm performance. Rather not using
	// this close can be referred in other function calls
	defer res.Body.Close()
	robots, err := ioutil.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%s", robots)

	// defer will take the argument at the time 
	// it is used, so this defer will actually print out 
	// "start"
	a9 := "start"
	defer fmt.Println(a9)
	a9 = "end"

	// Panic
	// no exception in Go
	// usually jus throw an error value
	// unless the program cannot continue

	// a7, b7 := 1, 0
	// ans := a7 / b7
	// fmt.Println(ans)

	// panic("somthing bad happened")

	// panic happens after defer happens
	a9 = "start"
	defer fmt.Println(a9)
	// panic("something bad happened")
	defer func() {
		// recover() will return nil is it is not panic
		// return the error statement if it is 
		if err:= recover(); err!=nil {
			log.Println("Err", err)
		}
	}() // the paren is invoking the function
	// panic("somthing bad happened")
	// you need to panic() to pass to higher call stack
	// in order to handle it 
	panicker() 
	fmt.Println("end")

	// Pointers
	// copy value that the variable points to
	var aval int = 42
	var bpointer *int = &aval // declere pointer type
	fmt.Println(aval, *bpointer) // dereferencing

	// pointer arithmetic is not supported
	// unless you use the "unsafe" package

	// new function
	var ms *myStruct
	ms = &myStruct{foo: 24}
	(*ms).foo = 29
	// or just no paren
	ms.foo = 89
	ms = new(myStruct)
	fmt.Println(ms)
		
	// if not initialize, then it becomes nil
	// array is copy by value, but slice is not
	// difference is that one is [] the other is [3]
	// also map. It will point to the same object
}

type myStruct struct {
	foo int
}

func panicker() {
	fmt.Println("start")
	defer func() {
		if err := recover(); err != nil {
			log.Println("error:", err)
			// unless you re-panic the err
			panic(err)
		}
	}()
	// this is handled, but up in the call stack
	// it will continue
	panic("something bad happened")
}