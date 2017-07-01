package main

import "fmt"

func main(){
    var x [5]int;
    fmt.Println(x);
    x[3], x[2] = 12, 23
    fmt.Println(x);
    fmt.Println("Length of the given array :", len(x));
}
