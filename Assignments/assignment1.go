// Profit calculator
// It should take taxRate, Revenue, Expenses from user.
// EBT = R - E
// P = EBT * (1-taxRate/100)
package main

import (
	"fmt"
)

func main() {
	var revenue, expense, taxRate float64
	fmt.Println("Enter Revenue")
	fmt.Scan(&revenue)

	fmt.Println("Enter Expense")
	fmt.Scan(&expense)

	fmt.Println("Enter Taxrate")
	fmt.Scan(&taxRate)

	ebt := revenue - expense
	fmt.Println("Total Earning Before The Tax Is :", ebt)

	profit := ebt * (1 - taxRate/100)
	fmt.Println("Profit You Earned :", profit)

}
