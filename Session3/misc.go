package main

import "fmt"

func main() {
	defer fmt.Println("Nandkumar")
	fmt.Println("Hello")
	fmt.Println("Apppp")

	fmt.Println(fmt.Errorf("some random error message"))

	panic("Time To Panic")
}
