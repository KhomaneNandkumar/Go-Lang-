/*  ### Fizzbuzz
- User will provide a number and you need to prints Fizz, Buzz, FizzBuzz accrodingly.
   - Fizz => number is divisible by 3
   - Buzz => number is divisible by 5
   - Fizzbuzz => number is divisible by both 3 and 5
*/

package main

import "fmt"

func main() {
	for {
		var number int
		fmt.Println("Enter Number")
		fmt.Scan(&number)

		if number%3 == 0 && number%5 == 0 {
			fmt.Println("Fizzbuzz")
		} else if number%3 == 0 {
			fmt.Println("Fizz")
		} else if number%5 == 0 {
			fmt.Println("Buzz")
		} else {
			continue
		}
	}
}
