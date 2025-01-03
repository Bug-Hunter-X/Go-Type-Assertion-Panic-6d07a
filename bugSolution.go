```go
package main

import "fmt"

func main() {

    var i interface{}
    i = 10
    j := i.(int)
    fmt.Println(j * 2)

    i = "hello"
    switch i := i.(type) {
    case int:
        fmt.Println(i * 2)
    case string:
        fmt.Println(i)
    default:
        fmt.Println("Unknown type")
    }
}```