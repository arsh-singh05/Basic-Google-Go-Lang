package main

import "fmt"

type Node struct {
    value int
    next  *Node
}

func main() {
    head := &Node{value: 1}
    second := &Node{value: 2}
    third := &Node{value: 3}

    head.next = second
    second.next = third

    for current := head; current != nil; current = current.next {
        fmt.Println(current.value)
    }
    // prints 1, 2, 3
}

/*
This program uses pointers to create a linked list of Node structs. Each Node struct has a value field and a next field, which is a pointer to the next Node in the list.
The program uses the next pointers to link the nodes together and then iterates over the list using a for loop to print the values of each node.
*/
