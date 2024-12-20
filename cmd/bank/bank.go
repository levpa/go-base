package main

import "fmt"

func main() {

	var accountBalance = 1000.0

	fmt.Println("Welcome to Go Bank!")
	for {
			fmt.Println("What do you want to do?")
			fmt.Println("1. Check balance")
			fmt.Println("2. Deposit money")
			fmt.Println("3. Withdraw money")
			fmt.Println("4. Exit")

			var choice int
			fmt.Print("Your choice: ")
			fmt.Scan(&choice)

			// wantsCheckBalance := choice == 1

			if  choice == 1 {
				fmt.Println("Your balance is ", accountBalance)
			} else if choice == 2 {
				fmt.Print("How much do you want to deposit? ")
				var depositAmount float64; fmt.Scan(&depositAmount)

				if depositAmount <= 0 {
					fmt.Println("Invalid amount. Must be greater than 0.")
					continue
				}
				accountBalance += depositAmount
				fmt.Println("Balance updated! New amount: ", accountBalance)
			}
	}

}
