package main

import "fmt"

func method1() {
    fmt.Println("Executing method1")
}

func method2() {
    fmt.Println("Executing method2")
}

func method3() {
    fmt.Println("Executing method3")
}

func main() {
    defer method3()
    defer method2()
    defer method1()
    fmt.Println("hello")
}
