package main

import (
	"fmt"
	"os"
)

func writeToFile(f string, d []byte) {
	err := os.WriteFile(f, d, 0644)
	if err != nil {
		fmt.Printf("Error:%v\n", err)
	}
}

func readFromFile(f string) {
	d, err := os.ReadFile(f)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
	}
	fmt.Println(string(d))
}

func main() {
	const filename = "output.txt"

	var input string
	fmt.Scan(&input)

	writeToFile(filename, []byte(input))
	readFromFile(filename)

}
