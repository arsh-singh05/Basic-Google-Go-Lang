package main

import "fmt"

// function that returns another function
func adderFactory(x int) func(int) int {
    return func(y int) int {
        return x + y
    }
}

func main() {
    add5 := adderFactory(5)
    add10 := adderFactory(10)
    fmt.Println(add5(3))
    fmt.Println(add10(3))
}

/*
This program demonstrates how a function can return another function. 
The adderFactory function returns a closure function that adds a value to its input.
The returned function is assigned to a variable and can be called multiple times with different input values.
*/
