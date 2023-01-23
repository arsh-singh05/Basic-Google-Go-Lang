package main

import "fmt"

func add(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

func main() {
    add5 := add(5)
    fmt.Println(add5(3)) // 8
}

/*
This program demonstrates how to use currying in Go, where a function that takes multiple arguments is transformed into a chain of functions each taking a single argument.
In this example, the add function is curried, so it can be called with the first argument, and it returns a function that takes the second argument.
*/
