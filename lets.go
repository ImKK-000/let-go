package main

import "fmt"

func say(message string) bool {
    return (len(message) > 0)
}

func main() {
    status := say("Welcome to Go Language!")
    fmt.Println(status)
}
