package main

import "fmt"

// function that returns a closure function
func adder() func(int) int {
    sum := 0
    return func(x int) int {
        sum += x
        return sum
    }
}

func main() {
    add := adder()
    for i := 0; i < 10; i++ {
        fmt.Println(add(i))
    }
}

/*
This program uses a closure function to create a simple adder.
The closure function is returned by the outer function, and it uses the variables declared in the outer function, but with different values in each call.
*/
