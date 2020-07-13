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
  




