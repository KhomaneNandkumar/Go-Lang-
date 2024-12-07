package main

import "fmt"

func main() {
	var actor map[string]int = map[string]int{
		"amir":        50,
		"srk":         51,
		"salomanbhai": 52,
	}
	fmt.Println(actor)
	fmt.Println(actor["srk"])
	fmt.Println(actor["srk3"])
	fmt.Println(actor)
}
