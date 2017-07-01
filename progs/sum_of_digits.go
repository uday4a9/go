package main

import "fmt"

func main(){
    var input int64;
    fmt.Print("Enter a integer: ");
    fmt.Scanf("%d", &input);

//    fmt.Println("Given Value is:", input);

    num := input;

    var res int64;
    res = 0;

    for ;num!=0; {
        res += (num%10);
        num /= 10;
    }
    fmt.Println("Sum of digits of",  input, "is:", res);
}
