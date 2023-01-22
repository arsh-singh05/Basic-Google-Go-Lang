package main

import "fmt"

type error interface {
    Error() string
}

type SyntaxError struct {
    line int
    msg  string
}

func (e SyntaxError) Error() string {
    return fmt.Sprintf("syntax error on line %d: %s", e.line, e.msg)
}

type ValueError struct {
    value int
    msg   string
}

func (e ValueError) Error() string {
    return fmt.Sprintf("invalid value %d: %s", e.value, e.msg)
}

func checkError(e error) {
    switch e.(type) {
    case SyntaxError:
        fmt.Println("Handling syntax error:", e)
    case ValueError:
        fmt.Println("Handling value error:", e)
    default:
        fmt.Println("Unknown error:", e)
    }
}

func main() {
    checkError(SyntaxError{line: 10, msg: "missing semicolon"})
    checkError(ValueError{value: 5, msg: "value out of range"})
}

/*
This program defines two structs SyntaxError and ValueError that implement the error interface by implementing the Error() method.
The program also defines a function checkError that takes an argument of type error and uses a type switch to handle different types of errors, in this case SyntaxError and ValueError.
The function main calls thecheckErrorfunction with two different types of errors,SyntaxErrorandValueError`, and the function handles them accordingly.
*/
