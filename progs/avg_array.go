package main

import "fmt"


func method1(){
    var x [5]int64;
    x[0] = 123
    x[1] = 234
    x[2] = 45
    x[3] = 56
    x[4] = 67

    fmt.Println("Array :", x)

    var total int64;
    for i:=0; i<5; i++ {
        total += x[i]
    }

    fmt.Println("Average of", x, "is", total / int64(len(x)))
}

func method2(){
    var x [5]int64;
    x[0] = 123
    x[1] = 234
    x[2] = 45
    x[3] = 56
    x[4] = 67

    fmt.Println("Array :", x)

    var total int64;
    for i:=0; i<len(x); i++ {
        total += x[i]
    }

    fmt.Println("Average of", x, "is", total / int64(len(x)))
}

func method3() {
    x:= [5]int64{123, 234, 345, 456, 567};

    for i, value := range x{
        fmt.Print(i, value, " ");
    }
    fmt.Println("");
}

func method4() {
    x:= [5]int64{123, 234, 345, 456, 567};
    y := x[1:3];
    fmt.Println("y", y);
    z := append(y, 23, 34);
    fmt.Println("z", z);
    fmt.Println("y", y);
}

func main(){
    method1();
    method2();
    method3();
    method4();
}
