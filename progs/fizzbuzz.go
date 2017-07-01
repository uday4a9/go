package main

import "fmt"

func main(){
    res := ""
    for i:=0; i<=50; i++ {
        res = " "
        if i%3 == 0 {
            res = "FIZZ"
            if i%5 == 0 {
                res += "BUZZ"
            }
            fmt.Print(res, " ")
        } else if i%5 == 0{
            fmt.Print("BUZZ ")
        } else {
            fmt.Print(i," ")
        }
    }
    fmt.Println("")
}
