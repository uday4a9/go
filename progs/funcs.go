package main

import ("fmt"
        "os")

func add(a int64, b int64) int64{
    return a + b;
}

func sub(a int64, b int64) int64 {
    return 123;
}

func mul(a int64, b int64) int64 {
    return 123;
}

func div(a int64, b int64) int64 {
    return 123;
}

func main(){
    fmt.Println("Calculator **********");
    var choice int8 = -1
    var a, b int64;

    fmt.Println(a, b, "Values")
    for ;; {
        fmt.Print("0. Exit\n1. Add\n2. Sub\n3. Mul\n4. Div\n");
        fmt.Print("Enter your choice from above listed: ")
        fmt.Scanf("%d", &choice);
        if choice != 0 && choice != -1{
            fmt.Print("Enter two numbers : ");
            fmt.Scanf("%d\n%d", &a, &b);
        }
        switch(choice){
            case 0: 
                    fmt.Println("Exiting form the program")
                    os.Exit(0)
            case 1: res:=add(a, b)
                    fmt.Println("Result:", res)
            case 2: res:=sub(a, b)
                    fmt.Println("Result:", res)
            case 3: res:=mul(a, b)
                    fmt.Println("Result:", res)
            case 4: res:=div(a, b)
                    fmt.Println("Result:", res)
        }
    }

}
