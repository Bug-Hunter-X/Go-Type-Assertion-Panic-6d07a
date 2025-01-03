```go
package main

import "fmt"

func main() {

    var i interface{}
    i = 10
    j := i.(int)
    fmt.Println(j * 2)

    i = "hello"
    // This will cause a panic
    // j := i.(int) // this is wrong
    k := i.(string)
    fmt.Println(k)
}