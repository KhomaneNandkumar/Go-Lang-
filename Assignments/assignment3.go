/* ### Simple Calculator
  - Addition 1
  - Substraction 2
  - Division 3
  - Multiplication
  - Pow

    INPUT
	1
	3 4
    OUTPUT
	7

    INPUT
	2
	10 2
    OUTPUT
	8
*/

package main

import (
	"fmt"
	"math"
)

func addition(a, b float64) float64 {
	return a + b
}

func substraction(a, b float64) float64 {
	return a - b
}
func multiplication(a, b float64) float64 {
	return a * b
}
func division(a, b float64) float64 {
	return a / b
}
func power(a, b float64) float64 {
	return math.Pow(a, b)
}

func main() {
	var choice int
	var number1, number2 float64
	for {
		fmt.Printf("Enter Your Choice \n1)For Addition: \n2)For Substraction: \n3)For Division: \n4)For Multiplication: \n5)For Power: \n6) For Exit: \n")
		fmt.Scan(&choice)

		if choice == 6 {
			break
		}

		fmt.Println("Enter First Number")
		fmt.Scan(&number1)

		fmt.Println("Enter Second Number")
		fmt.Scan(&number2)

		switch choice {
		case 1:
			fmt.Println("Addition Of ", number1, "and", number2, "is :", addition(number1, number2))

		case 2:
			fmt.Println("Substraction Of ", number1, "and", number2, "is :", substraction(number1, number2))

		case 3:
			fmt.Println("The Division Of ", number1, "and", number2, "is :", division(number1, number2))

		case 4:
			fmt.Println("Multiplication Of ", number1, "and", number2, "is :", multiplication(number1, number2))

		case 5:
			fmt.Println("The Power Of ", number1, "Raised To", number2, "Is :", power(number1, number2))

		default:
			fmt.Println("Please Enter Valid Choice")

		}
	}
}
