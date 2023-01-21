package main

import 
(
    "fmt"
    "net/http"
)

func handler(w http.ResponseWriter, r *http.Request) 
{
    fmt.Fprintf(w, "Hello, World!")
}

func main() 
{
    http.HandleFunc("/", handler)
    http.ListenAndServe(":8080", nil)
}

/*
program uses the net/http package to create a simple HTTP server.
The handler function is registered as the handler for the root path / of the server, it writes "Hello, World!" to the response writer.
The ListenAndServe function starts the server and listens on port 8080, any incoming request to the server will be handled by the registered handler.
/*
