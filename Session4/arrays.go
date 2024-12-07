package main

import "fmt"

func main() {
	var arr = []int{1, 2, 3, 4, 5}
	fmt.Println(arr)

	fmt.Println(append(arr, 6))
	fmt.Println(arr)
	fmt.Println(arr[2])
	fmt.Println(arr[:3])
	fmt.Println(arr[1:3])
	fmt.Println(arr[3:])
	fmt.Println(arr)
}
