package main

import ("fmt"
        "strings"
    )

func main() {
    s := "Hello world"
    fmt.Println(strings.ToUpper(s))
    fmt.Println(len(strings.Split(s, " ")))
}
