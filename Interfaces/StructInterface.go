package main

import "fmt"

type Notifier interface {
    notify()
}

type User struct {
    name  string
    email string
    notifier Notifier
}

type EmailNotifier struct {}

func (e EmailNotifier) notify() {
    fmt.Println("Sending email to", e.email)
}

func main() {
    user := User{
        name: "John Doe",
        email: "johndoe@example.com",
        notifier: EmailNotifier{},
    }
    user.notifier.notify()
}

/*
This program defines an interface Notifier with a single method notify().
A struct User is defined that has a field of type Notifier. 
The struct EmailNotifier implements the Notifier interface by implementing the notify() method.
The User struct can be used to define a user with a specific email notifier, in this case it is EmailNotifier.
*/
