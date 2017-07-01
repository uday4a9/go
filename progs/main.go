package main

import "fmt"


func arithmetic(){
    fmt.Println("\n******* Below are int operations ********")
    fmt.Println("int addition    : 2 + 1 =", 2 + 1);
    fmt.Println("int subtraction : 2 - 1 =", 2 - 1);
    fmt.Println("int multiplication : 2 * 3 =", 2 * 3);
    fmt.Println("int division : 2 / 3 =", 2 / 3);
    fmt.Println("int reaminder: 2 % 3 =", 2 % 3);
    fmt.Println("\n******* Below are float operations ********")
    fmt.Println("float addition    : 2 + 1 =", 2 + 1);
    fmt.Println("float subtraction : 2 - .5 =", 2.0 - .5);
    fmt.Println("float multiplication : 2.0 * 3.1 =", 2.0 * 3.1);
    fmt.Println("float division : 2 / 3.0 =", 2 / 3.0);
    fmt.Println("float division : 2.0 / 3 =", 2.0 / 3);
//    fmt.Println("float remainder : 2 % 3.0 =", 2 % 3.0);
}


func string_ops(){
    fmt.Println("\n******** Below are strin manipulations ******");
    fmt.Println("Hey.... are \n u der");
    // Raw string implementaion
    fmt.Println(`Hey.... are \n u der`);
    fmt.Println("legth :", len("Hello world"));
    fmt.Println("Hello world"[1]);
    fmt.Println("Hellllo," + "World")
    fmt.Println("Hellllo," + `World`)
    fmt.Println("Hellllo,", `World`)
}

func booleans(){
    fmt.Println("\n******** Below are booleans *********");
    fmt.Println(true);
    fmt.Println(false);
    fmt.Println(true && false);
    fmt.Println(true && true);
    fmt.Println(false && true);
    fmt.Println(false && false);
    fmt.Println(true || false);
    fmt.Println(true || true);
    fmt.Println(false || true);
    fmt.Println(false || false);
    fmt.Println(!false);
    fmt.Println(!true);

}

func main(){
    fmt.Println("Hello world, welcome to GO");
    arithmetic();
    string_ops();
    booleans();
}
