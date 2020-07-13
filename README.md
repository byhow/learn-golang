## Golang

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

#### Variables
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
  

#### Primitive Types
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