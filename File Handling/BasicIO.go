package main

import 
(
    "fmt"
    "os"
)

func main() 
{
    file, err := os.Create("test.txt")
    if err != nil 
    {
        fmt.Println(err)
        return
    }
    defer file.Close()
    file.WriteString("Hello, World!")
}

/*
This program uses the os package to create a new file named test.txt and write the string "Hello, World!" to it.
The Create function returns a file handle and an error, it checks the error and if it's not nil, it prints the error and exits the program.
The defer statement makes sure that the file handle is closed after the program is finished executing.
*/
