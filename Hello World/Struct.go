package main

import "fmt"

type Person struct 
{
    name string
    age int
}

func main() 
{
    p := Person{"John", 25}
    fmt.Println(p.name, p.age)
}

// This program declares a struct Person that has two fields name and age of type string and int respectively, then it creates a variable p of type Person and assigns values to its fields.
// Then it uses the fmt package to print the name and age fields of the struct.