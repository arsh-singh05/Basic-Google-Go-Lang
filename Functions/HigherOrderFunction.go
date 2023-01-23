package main

import "fmt"

// Higher-order function
func apply(n int, f func(int) int) int {
    return f(n)
}

func main() {
    square := func(n int) int {
        return n * n
    }
    fmt.Println(apply(5, square))
}

/*
This program demonstrates how to use a higher-order function, which is a function that takes another function as a parameter or returns a function as a result. 
In this example, the apply function takes an integer and a function as parameters and applies the function to the integer.
*/
