package main

import (
	"bufio"
	"fmt"
	"multimiser/dice"
	"multimiser/evenOrOdd"
	standaradCalculator "multimiser/standardCalculator"
	"os"
	"strconv"
)

func main() {

	fmt.Println("Salute! This is multimiser")
	fmt.Println("This program allows you to utilize the next options:")
	fmt.Println("Choose one of them \n")
	for {
		for {
			fmt.Println("1) standard calculator	( press 1 )")
			fmt.Println("2) Even or Odd checker	( press 2 )")
			fmt.Println("3) Dice			( press 3 )")
			fmt.Println("4) Texas holdem		( press 4 )")
			fmt.Println("5) Exit the program	( press 5 or 0)")
			fmt.Println("\n>>> Main menu <<<\n")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Println("It's not an integer number! type the numbers from the list")
				break
			}

			switch input {
			case 1:
				standaradCalculator.StandardCalculator()
			case 2:
				evenOrOdd.EvenOrOdd()
			case 3:
				dice.Dice()
			case 4:
				fmt.Println("This option is in development\n")
			case 5, 0:
				fmt.Println("The end of current session")
				return
			default:
				fmt.Println("This option is not available")
			}
		}
	}
}
