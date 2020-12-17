package main

import "fmt"

func HelloWorld() string {
	return "Hello world"
}

func main() {
	msg := HelloWorld()
	fmt.Println(msg)
}
