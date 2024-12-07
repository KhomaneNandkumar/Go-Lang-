package main

import (
	"fmt"
	"time"
)

type Student struct{
	username  string
	age  string
	gender  string
	id int
	joinedAt time.Time
}

func getData(prompt string) string {
	fmt.Println(prompt)
	var input string

	fmt.Scan(&input)
	return input
}

func outputData(name, age, gender string) {
	fmt.Println("Name is", name)
	fmt.Println("Agfe Is", age)
	fmt.Println("Genser is", gender)
}
func main() {
	name := getData("Enter Name:")
	age := getData("Enter Age:")
	gender := getData("Enter Gender:")

	var studentA Student{
		username : name ,
		age : age, 
		gender : gender,
		id :1,
		joinedAt time
	}

	outputData(name, age, gender)

}
