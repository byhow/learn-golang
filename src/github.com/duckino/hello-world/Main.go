package main

import (
	"fmt"
	"strconv"
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
}
