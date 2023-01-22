package main

import (
    "fmt"
    "os"
)

func main() {
    file, err := os.Open("example.txt")
    if err != nil {
        fmt.Println(err)
        return
    }
    defer file.Close()

    data := make([]byte, 100)
    count, err := file.Read(data)
    if err != nil {
        fmt.Println(err)
        return
    }
    fmt.Println(count, string(data[:count]))
}
