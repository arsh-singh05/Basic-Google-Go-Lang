package main

import 
(
    "fmt"
    "time"
)

func printNumbers() 
{
    for i := 1; i <= 5; i++ 
    {
        time.Sleep(1 * time.Second)
        fmt.Printf("%d ", i)
    }
}

func printLetters() 
{
    for i := 'a'; i <= 'e'; i++ 
    {
        time.Sleep(1 * time.Second)
        fmt.Printf("%c ", i)
    }
}

func main() 
{
    go printNumbers()
    go printLetters()
    time.Sleep(5 * time.Second)
}

/*
This program uses the go keyword to create two goroutines (lightweight threads of execution), printNumbers and printLetters that run concurrently.
The main function waits for 5 seconds before exiting, to give the goroutines time to complete.
/*
