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
  

### Loops 
* For loop
  * no while keyword
  * simple loop
    * for initializer; test; incrementor {}
    * for test {}
    * for {} // infinite loop
  * exiting early
    * break
    * continue
    * labels
  * loop over collections
    * arrays, slices, maps, string, channels
    * for k, v := range collection {}

### Defer, Panic, Recover
* Defer
  * Used to delay execution of a statement until function exits
  * Useful to group "open" and "close" functions together
    * Be careful in loops
  * Run in LIFO (last-in, first-out) order
  * Arguments evaluated at time defer is executed, not at time of called function execution
* Panic
  * Occur when program cannot continue at all
    * Don't use when file cannot be opened, unless it is critical
    * Use for unrecoverable events - cannot obtain TCP port for web server
  * Function will stop executing
    * Deferred functions will still fire
  * If nothing handles panic, program will exit
* Recover
  * Used to recover from panics
  * Only useful in deferred functions
  * Current function will not attempt to continue, but higher functions in call stack will


### Pointers
* Create pointers
  * *int - a pointer to an integer
  * Use the addressof operator (&) to get address of variable
* Dereferencing pointers
  * Dereference pointers by preceding with an asterisk(*)
  * Complex types (e.g. struct) are automatically dereferenced
* Create pointers to objects
  * Can use the addressof operator (&) if value types already exists
    * ms := myStruct{foo: 42}
    * p := &ms
  * Use addressof operator before initializer
    * &myStruct {foo: 42}
  * Use the new keyword
    * Can't initialize fields at the same time
* Types with internal pointers
  * All assignment operations in Go are copy operations
  * Slices and maps contain internal pointers, so copies point to same underlying data


### Functions
* Basic Syntax
  * ```func foo() {}```
* Parameters
  * Comma delimited list of variables and types
    * ```func foo(bar string, bar int)```
  * Parameters of same type list type once
    * ```func foo(bar, baz int)```
  * When pointers are passed in, the function can change the value in the caller
    * This is always true for data of slices and maps
  * Use variadic parameters to send list of same types in
    * Must be last parameter
    * Received as a slice
    * func foo (bar string, baz ...int)
* Return values
  * Single return values just list type
    * `func foo() int`
  * Multiple return value list types surrounded by parentheses
    * `func foo() (int, error)`
    * The (result type, error) paradigm is a very common idiom
  * Can use named return values
    * initialize returned variables
    * Return using return keyword on its own
  * Can return addresses of local variables
    * Automatically promoted from local memory stack to shared memory heap
* Anonymous functions
  * Functions don't have names if they are:
    * Immediately invoked
      * func() {...} ()
    * Assigned to a variable or passed as an argument to a function:
      * a := func() {} a()
* Functions as types
  * Can assign functions to variables or use as arguments and return values in functions
  * Type signature is like function signature, with no parameter names
    * `var f func(string, string, int) (int, error)
* Methods
  * Function that executes in context of a type
  * Format
    * `func (g greeter) greet() {...}`
  * Receiver can be value or pointer
    * Value receiver gets copy of type
    * Pointer receiver gets pointer to type


### Interfaces 
* Best Practices
  * Use many, small interfaces
    * Single method interfaces are some of the most powerful and flexible
      * `io.Writer, io.Reader, interface{}`
  * Don't export interfaces for types that will be consumed
  * Do export the interfaces for types that will be used by package
  * Design functions and methods to receive interfaces whenever possible

* Basics
```go
type Writer interface {
  Write([]byte) (int, error)
}
type ConsoleWriter struct {}

func (cw ConsoleWriter) Write(data []byte) (int, error) {
  n, err := fmt.Println(string(data))
  return n, err
}
```
* Composing interfaces
```go
type Writer interface {
  Write([]byte) (int, error)
}
type Closer interface {
  Close() error
}
type WriterCloser interface {
  Writer
  Closer
}
```

* Type conversion
```go
var wc WriterCloser = NewBufferedWriterCloser()
bwc := wc.(*BufferedWriterCloser)
```

* The empty interface and type switches
```go
var i interface{} = 0
switch i.(type) {
  case int:
    fmt.Println("i is an integer")
  case string:
    fmt.Println("i is an string")
  default:
    fmt.Println("i dont know what i is")
}
```

* Implementing with values vs pointers
  * Method set of value is all methods with value receivers
  * Method set of pointer is all methods, regardless of receiver type

* Best Practices
  * Don't export interfaces for types that will be consumed


### Goroutine
* Best practices
  * Don't create goroutine in libraries
    * Let consumer control concurrency
  * When creating a goroutine, know how it will end
    * Avoids subtle memory leaks 
  * Check for race conditions at compile time
    * `go run -race <args>`
    * recommended to run before prod
* Creating goroutine
  * Use `go` keyword in front of function call
  * When using anonymous functions, pass data as local variables
* Synchronization
  * Use sync.waitGroup
  * Use sync.Mutex and sync.RWMutex to protect data access
* Parallelism
  * By default, Go will use CPU threads equal to available cores
  * Change with runtime.GOMAXPROCS
  * More threads can increase performance, but too many can slow it down