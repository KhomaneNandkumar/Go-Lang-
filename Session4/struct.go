package main

import (
	"fmt"
	"time"
)

func getData(prompt string) string {
	fmt.Println(prompt)
	var input string

	fmt.Scan(&input)
	return input
}

type User struct {
	username   string
	age        string
	gender     string
	lastActive time.Time
}

func main() {
	var dummyUser User

	name := getData("Enter Name:")
	age := getData("Enter Date Of Birth:")
	gender := getData("Enter Gender:")

	dummyUser = User{
		username:   name,
		age:        age,
		gender:     gender,
		lastActive: time.Now(),
	}
	displayDataAdv(dummyUser)
}

func displayDataAdv(userData User) {
	fmt.Println("Name Is :", userData.username)
	fmt.Println("Age Is :", userData.age)
	fmt.Println("Gender IS :", userData.gender)
	fmt.Println("Last Active At :", userData.lastActive)

}
