package main

import "fmt"

func main() {
	age := 32

	ageP := &age
	fmt.Println("Age Is:", *ageP) //dereference the address
	fmt.Println("Adult Age Version 1:", getAdultYearsUpdated(ageP))
	//fmt.Println("Adult Age Version 2:",getAdultYearsUpdated2(ageP))
}

func getAdultYears(age int) int {
	return age - 18
}

func getAdultYearsUpdated(ageP *int) int {
	return *ageP - 18
}

func getAdultYearsUpdated2(ageP *int) {
	*ageP = *ageP - 18

}
