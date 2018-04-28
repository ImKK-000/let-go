package main

import "fmt"

func say(message string) (bool, string) {
	fmt.Println(message)
	return true, message
}

func main() {
	status, originalText := say("yo")
	fmt.Println(status, originalText)
}
