package main

import (
    "fmt"
)

func recoverName() {
    if r := recover(); r != nil {
        fmt.Println("recovered from ", r)
    }
}

func fullName(firstName *string, lastName *string) {
    defer recoverName()
    if firstName == nil {
        panic("runtime error: first name cannot be nil")
    }
    if lastName == nil {
        panic("runtime error: last name cannot be nil")
    }
    fmt.Printf("%s %s\n", *firstName, *lastName)
    fmt.Println("returned normally from fullName")
}

func main() {
    defer fmt.Println("done")
    var firstName *string = new(string)
    *firstName = "Bob"
    var lastName *string = new(string)
    *lastName = "Smith"
    fullName(firstName, lastName)
    fullName(nil, lastName)
}

/*
In this example, the function "fullName" checks if the firstName and lastName are not nil, if not, it will panic with a message, otherwise it will print the fullName. 
The function recoverName is deferred and will be called when the function panics, it will recover from the panic, print the message and continue the execution of the program
*/
