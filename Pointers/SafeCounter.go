package main

import (
    "fmt"
    "sync"
)

type SafeCounter struct {
v map[string]int
mux sync.Mutex
}

func (c *SafeCounter) Inc(key string) {
c.mux.Lock()
c.v[key]++
c.mux.Unlock()
}

func (c *SafeCounter) Value(key string) int {
c.mux.Lock()
defer c.mux.Unlock()
return c.v[key]
}

func main() {
c := SafeCounter{v: make(map[string]int)}
for i := 0; i < 1000; i++ {
go c.Inc("somekey")
}
 fmt.Println(c.Value("somekey"))
}

/*
This program uses pointers to create a concurrent safe data structure, a counter that can be incremented by multiple goroutines simultaneously.
The program uses a struct `SafeCounter` which has a map `v` to store the counter value and a mutex `mux` to synchronize access to the shared counter.
The `Inc` method acquires the lock, increments the counter and releases the lock.
The `Value` method acquires the lock, reads the counter value and releases the lock.
*/
