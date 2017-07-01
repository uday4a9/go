package main

import "fmt"

func main() {
    fmt.Println("Hello World", 22/7.0)
    x := fun()
    fmt.Println("function return value:", x)
    foo()
    y := boo()
    fmt.Println("function return value:", y)

//    panic("Yes.... Implemented")

    rc1, rc2 := ret_multiple()
    fmt.Println("Return values: ", rc1, "and", rc2)

}

func fun() int {
    return 123
}

func foo() {
    fmt.Println("Yes... in foo method")
    return
}

func boo() (r int) {
    r =321
    return
}

func ret_multiple() (int, int) {
    return 21, 32
}
