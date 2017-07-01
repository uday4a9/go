package main

import "fmt"

func main() {
    i := 10
    j := 0
    fmt.Println("hello world", i, &i, *&i, &j)
}
