package main

import "fmt"

// function that takes another function as a parameter
func execute(f func(string)) {
    f("Hello World")
}

func main() {
    execute(fmt.Println)
}

/*
This program shows how to pass a function as a parameter to another function.
In this example, the execute function takes a function as a parameter and calls it with the argument "Hello World".
*/
