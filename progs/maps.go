package main

import "fmt"

func method(){
//    var hash map[string]int;
    hash := make(map[string]int)
    var str string;

    fmt.Print("Enter a string: ");
    fmt.Scanf("%s", &str);
    fmt.Println("GIven string :", str);

    for _, val:= range str{
//        fmt.Print(string(val), " ");
        hash[string(val)] += 1;
    }
    fmt.Println("hash", hash, "length is:", len(hash));

    for key, val := range hash{
        fmt.Println(key, val)
    }

    key, val := hash["z"]
    fmt.Println("Result :", key, val)

    if key, val := hash["z"]; val {
        fmt.Println(key, val);
    }
}

func main(){
    method();
}
