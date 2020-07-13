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


}
