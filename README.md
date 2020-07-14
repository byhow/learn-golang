# Golang

Reason why:
* Strong and statically typed (e.g. Java)
* Excellent community
* Key features:
  * Simplicity
  * Fast compile times
  * Garbage collected
  * Built-in concurrency
  * Compile to standalone binaries (only one version of dependencies)

Resource:
* golang.org
  * package libraries
* Gopher slack & Go Forum
* Golang Bridge
* play.golang.org

```go
package main

import {
  "fmt"
}

func main() {
  fmt.Println("Hello World!")
}
```

### Variables
* Variable declaration
  * var foo int
  * var foo int = 42
  * foo := 42
* Redeclaration and shadowing
  * Cannot redelare, but can shadow variables
  * All variables has to be used, otherwise triggering compiler error
* Visibility
  * Lowercase first letter for package scope
  * Uppercase first letter to export
  * No private scope, but can do block
* Naming convention
  * Pascal or camelCase
    * Capitalize acronyms (HTTP, URL)
  * As short as reasonable
    * Longer names for longer lives
* Type Convertion
  * DestinationType (variable)
  * Use strconv package for strings (or got unicode conversion)
  

### Primitive Types
* Boolean 
  * Values are true or false
  * Not an alias for other types (e.g. int) 
  * Zero value is false
* Numeric 
  * Integer
    * Signed integers
      * int types has varying size, but min 32 bits
      * 8 bit (int8) through 64 bit (int64)
    * Unsigned integers
      * 8 bit (byte and uint8) through 32 bit (uint32)
    * Arithmetic operations
      * Addition, subtraction, multiplication, division, remainder
    * Bitwise operations
      * And, or, xor, and not
    * Zero value is 0
    * Cannot mix types in same family (uint16 + uint32 = error)
  * Floating points
    * Follow IEEE-754 standard
    * Zero value is 0
    * 32 and 64 bit versions
    * Literal styles
      * Decimal (3.14)
      * Exponential (13e18 or 2e10)
      * Mixed (13.7e2)
    * Arithmetic operations
      * Addition, subtraction, multiplication, division
  * Complex numbers
    * Zero value is (0 + 0i)
    * 64 and 128 bit versions
    * Built-in functions
      * `complex` - make complex number from two floats
      * `real` - get real part as float
      * `imag` - get imaginary part as float
    * Arithmetic operations
      * Addition, subtraction, multiplication, division
* Text types
  * Strings
    * UTF-8
    * Immutable
    * Can be concatenated with plus(+) operator
    * Can be converted to []byte
  * Rune
    * UTF-32
    * Alias for int32
    * Special methods normally required to process
      * e.g. strings.Reader#ReadRune


### Constants
* Immutable, but can be shadowed
* Replaced by the compiler at compile time
  * Value must be calculable at compile time
* Naming convention
  * Named like variables
    * PascalCase for exported constants
    * camelCase for internal constants
* Typed constants
  * Works like immutable variables
  * Can interoperate only with same type
* Untyped constants
  * Can interoperate with similar type
* Enumerated constants
  * Special symbol `iota` allows related constants to be created easily
  * `iota` at 0 in each const block and increments by one
  * Watch out of constant values that match zero values for variables
* Enumeration expressions
  * Operations that can be determined at compile time are allowed
    * Arithmetic
    * Bitwise operations
    * Bitshifting


### Array and Slices
* Arrays
  * Creation
    * Collection of items with same type
    * Fixed size
    * Declaration styles
      * `a := [3]int{1, 2, 3}`
      * `a := [...]int{1, 2, 3}`
      * `var a [3]int`
    * Access via zero-based index
      * `a := [3]int {1, 3, 5} // a[1] == 3`
    * `len` function returns size of array
    * copies refer to different size of array
    * Copies refer to different underlying data
  * Built-in functions
  * Working with arrays
* Slices
  * Backed by array
  * Creation styles
    * Slice existing array or slice
    * Literal style
    * Via make function
      * `a := make([]int , 10) // create slice with capacity and length == 10`
      * `a := make([]int, 10, 100) // slice with length == 10 and capacity == 100`
  * `len` function returns length of slice
  * `cap` function returns length of underlying array
  * `append` function to add elements to slice
    * May cause expensive copy operation if underlying array is too small
  * Copies refer to same underlying array

### Map and Struct
* Map (reference type)
  * Collections of value types that are accessed via keys
  * Created via literals or via make function
  * Members accessed via `[key]` syntax
    * `myMap["key"] = "value"`
  * Check for presence with "value, ok" form of result
  * Multiple assignments refer to same underlying data
* Struct
  * Collections of disparate data types that describe a single concept
  * Keyed by named fields
  * Normally created as types, but anonymous structs are allowed
  * Structs are value types
  * No inheritance, but can use composition via embedding
  * Tags can be added to struct field to describe field

### Control Flows
* If statements
  * Initializer
  * Comparison operators
  * Logical operators
  * Short circuiting
  * If - else statements
  * If - else if statements
  * Equality and floats (divide them and see if it is in a range)
* Switch statements
  * Switching on a tag
  * Cases with multiple tests
  * Initializers
  * Switch with no tag
  * Fallthrough
  * Type switches
  * Breaking out early
  



