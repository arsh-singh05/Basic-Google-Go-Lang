package main

import (
    "fmt"
    "sync"
)

func initialize() {
    fmt.Println("Initializing...")
}

var once sync.Once

func main() {
    once.Do(initialize)
    once.Do(initialize)
    once.Do(initialize)
}

/*
In this example, the "sync.Once" struct is used to ensure that the "initialize" function is called only once, no matter how many times the "once.Do(initialize)" function call is invoked.
*/
