package main

import "fmt"

type operation func(int, int) int

func add(a, b int) int {
    return a + b
}

func sub(a, b int) int {
    return a - b
}

func calculator(a, b int, op operation) int {
    return op(a, b)
}

func main() {
    a, b := 5, 3
    fmt.Println(calculator(a, b, add)) // 8
    fmt.Println(calculator(a, b, sub)) // 2
}

/*
This program demonstrates how to use an alias for a function type and use it as a parameter for another function. 
In this example, the operation type is created as an alias for a function type that takes two int parameters and returns an int.This type is then used as the type of the op parameter in the calculator function which takes 2 integers and an operation and returns the result of the operation.
*/
