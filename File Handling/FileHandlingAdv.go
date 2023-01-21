package main

import (
    "fmt"
    "io/ioutil"
    "os"
)

func main() 

    // Reading from a file
    data, err := ioutil.ReadFile("test.txt")
    if err != nil 
	{
        fmt.Println(err)
        return
    }

    // Writing to a new file
    err = ioutil.WriteFile("test_copy.txt", data, 0644)
    if err != nil 
	{
        fmt.Println(err)
        return
    }

    // Appending to an existing file
    file, err := os.OpenFile("test_copy.txt", os.O_APPEND|os.O_WRONLY, 0644)
    if err != nil 
	{
        fmt.Println(err)
        return
    }
    defer file.Close()

    _, err = file.WriteString("\nAppended text")
    if err != nil 
	{
        fmt.Println(err)
        return
    }

    fmt.Println("File copied and appended successfully")
}

/*
This program reads the contents of a file named "test.txt" using the ioutil.ReadFile function, then it writes the contents to a new file named "test_copy.txt" using the ioutil.WriteFile function.
Finally, it opens the "test_copy.txt" file in append mode and writes "Appended text" to the end of the file using the os.OpenFile and WriteString methods.
If there's an error in any of the steps, it will print the error and exit.
/*