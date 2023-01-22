package main

import (
    "fmt"
    "sync"
)

func printNumbers(wg *sync.WaitGroup) {
    for i := 1; i <= 5; i++ {
        fmt.Printf("%d ", i)
    }
    wg.Done() // signals the WaitGroup that this goroutine is done
}

func main() {
    var wg sync.WaitGroup
    wg.Add(1) // adds 1 to the WaitGroup counter
    go printNumbers(&wg) // starts the goroutine
    wg.Wait() // blocks until the WaitGroup counter is zero
    fmt.Println("main function")
}

/*
This program uses a sync.WaitGroup to wait for the printNumbers goroutine to complete before the main function exits.
The Add method is used to add 1 to the WaitGroup counter, and the Done method is used to signal that the goroutine has completed.
The Wait method is used to block execution of the main function until the WaitGroup counter is zero, indicating that all goroutines have completed.
/*
