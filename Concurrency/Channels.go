package main
import "fmt"

func sum(s []int, c chan int) {
    sum := 0
    for _, v := range s {
        sum += v
    }
    c <- sum // sends sum to channel c
}

func main() {
    s := []int{1, 2, 3, 4, 5}
    c := make(chan int)
    go sum(s[:len(s)/2], c)
    go sum(s[len(s)/2:], c)
    x, y := <-c, <-c // receives from channel c
    fmt.Println(x, y, x+y)
}

/*
This program uses two goroutines to calculate the sum of two slices of an array concurrently and then it uses a channel to synchronize the goroutines and receive the calculated sum from the goroutines.
The make(chan int) creates a new channel of type int.
The <-c operator is used to send and receive values from the channel. In this case, c <- sum sends the value of the sum variable to the channel, and x, y := <-c, <-c receives the values from the channel and assigns them to the variables x and y. 
This program demonstrates how channels can be used to synchronize and communicate between goroutines.
/*