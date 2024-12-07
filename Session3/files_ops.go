package main

import (
	"bufio"
	"fmt"
	"os"
)

func writer(f string, d []byte) {
	err := os.WriteFile(f, d, 0000)
	if err != nil {
		fmt.Printf("Error writing to File: %v\n", err)
		return
	}
	fmt.Printf("Successfully Wrote To File %s\n", f)
}

func reader(fs string) {
	file, err := os.Open(fs)
	if err != nil {
		fmt.Printf("Error From Reading File %v\n", err)
		return
	}
	defer file.Close()
	fmt.Printf("Reading Contenet From File %s\n ", fs)
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		fmt.Printf(scanner.Text())
	}
	if err := scanner.Err(); err != nil {
		fmt.Printf("Error Reading File: %v\n", err)
	}
}

func main() {

	const f = "example.txt"
	data := "Hiii , Good Evening all Of You My Name Is Name Is Nandkumar Ankush Khomane \n"
	fmt.Println("GOOO")

	writer(f, []byte(data))
	reader(f)
}
