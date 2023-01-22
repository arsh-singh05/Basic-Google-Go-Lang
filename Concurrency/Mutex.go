package main

import (
    "fmt"
    "sync"
)

var counter int
var mutex sync.Mutex

func incrementCounter(wg *sync.WaitGroup) {
    defer wg.Done()
    for i := 0; i < 10000; i++ {
        mutex.Lock()
        counter++
        mutex.Unlock()
    }
}

func main() {
    var wg sync.WaitGroup
    wg.Add(2)
    go incrementCounter(&wg)
    go incrementCounter(&wg)
    wg.Wait()
    fmt.Println("Final counter value:", counter)
}

/*
This program uses two goroutines to increment the value of a shared variable counter concurrently, and it uses a sync.Mutex to synchronize access to the shared variable. 
The Lock method is used to acquire the lock and the Unlock method is used to release the lock, ensuring that only one goroutine can access the shared variable at a time.
This program demonstrates how mutex can be used to synchronize access to shared resources and prevent race conditions.
/*