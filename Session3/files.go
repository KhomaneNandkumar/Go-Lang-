package main

import (
	"fmt"
	"os"
)

func writeToFile(f string, d []byte) {
	os.WriteFile(f, d, 0644)

}

func readFromFile(f string) {
	d, _ := os.ReadFile(f)
	fmt.Println(string(d))

}

func main() {
	const filename = "app.txt"

	fmt.Println("Enter The Text You Want To Write In The File")
	var input string
	fmt.Scan(&input)

	writeToFile(filename, []byte(input))
	readFromFile(filename)
}
