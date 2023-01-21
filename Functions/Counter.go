package main

import "fmt"

func counter() func() int {
    count := 0
    return func() int {
        count++
        return count
    }
}

func main() {
    c := counter()
    fmt.Println(c()) // 1
    fmt.Println(c()) // 2
    fmt.Println(c()) // 3
}

/*
This program defines a function counter that returns a function.
The returned function increments and returns a variable count that is defined within the counter function.
The main function assigns the returned function to the variable c and calls it multiple times, each time it will increment the count and print the value of the count, thus it prints 1,2,3.
/*