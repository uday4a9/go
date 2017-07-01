package main

import ("fmt"
        "os"
    )

func main() {
    fmt.Println("WElcome to new file ops program")
    filereader("sim.txt")
}

func filereader(filename string) {
//    fmt.Println("*" * 20)
    fmt.Println("got the filename as", filename)
    file, err := os.Open(filename)
    defer file.Close()

    if err != nil {
        fmt.Println("got an error")
        return
    }
//    fmt.Println(file, err)

    fstat, err := file.Stat()
    if err!=nil {
        fmt.Println("Couldn't stat on the file")
        return
    }

    content := make([]byte, fstat.Size())

    _, err = file.Read(content)
    if err!=nil {
        fmt.Println("Couldn;t get what expected")
        return
    }

    print("File contents :", string(content))

//    fmt.Println("*" * 20)
}
