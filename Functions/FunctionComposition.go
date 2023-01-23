package main

import "fmt"

func add1(n int) int {
    return n + 1
}

func square(n int) int {
    return n * n
}

func main() {
    // function composition
    f := add1
    g := square
    h := func(x int) int { return f(g(x)) }
    fmt.Println(h(2)) // 9
}

/*
This program demonstrates how to compose functions, where the output of one function is passed as the input to another.
In this example, the h function is created by composing f(add1) and g(square), where h(x) = f(g(x)).
*/
