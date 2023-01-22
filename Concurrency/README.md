## Concurrency 
It is used in Go is used to write concurrent and parallel programs.Go has built-in support for goroutines and channels, which can be used to write concurrent and parallel programs.Go also has support for synchronization primitives such as WaitGroup and Mutex, which can be used to synchronize access to shared resources.

Concurrency is a fundamental concept in Go programming language and allows you to write programs that can perform multiple tasks simultaneously. Go provides built-in support for concurrency through the use of goroutines and channels.

A goroutine is a lightweight thread of execution that runs concurrently with other goroutines in the same program. Goroutines are created using the go keyword followed by a function call. For example, go someFunc() starts the execution of the someFunc function as a goroutine.

Channels are a way to communicate between goroutines. They can be used to send and receive values between goroutines. Channels are created using the make function, and the <- operator is used to send and receive values from a channel. For example, c := make(chan int) creates a channel of integers and c <- 1 sends the value 1 to the channel.

Concurrency can be a powerful tool for creating high-performance and concurrent programs in Go, but it also requires a deeper understanding of concurrency patterns and synchronization mechanisms like [WaitGroup](https://github.com/arsh-singh05/Basic-Google-Go-Lang/blob/main/Concurrency/WaitGroup.go) and [Mutex](https://github.com/arsh-singh05/Basic-Google-Go-Lang/blob/main/Concurrency/Mutex.go).

In summary, Go provides built-in support for concurrency through goroutines and channels, which can be used to write concurrent and parallel programs, however they also require a deeper understanding of concurrency patterns and synchronization mechanisms.
