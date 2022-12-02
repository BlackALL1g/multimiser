package main

import (
	"fmt"
	"multimiser/evenOrOdd"
	"multimiser/russianRoulette"
	standaradCalculator "multimiser/standardCalculator"
)

func main() {

	var op int

	fmt.Println("Salute! This is multimiser")
	fmt.Println("This program allows you to utilize the next options:")
	fmt.Println("Choose one of them \n")

	for {
		fmt.Println("1) standard calculator	( press 1 )")
		fmt.Println("2) Even or Odd checker	( press 2 )")
		fmt.Println("3) Russian Roulette	( press 3 )")
		fmt.Println("4) Texas holdem		( press 4 )")
		fmt.Println("5) Exit the program	( press 5 )")
		fmt.Println("\n>>> Main menu <<<\n")

		fmt.Scan(&op)
		switch op {
		case 1:
			standaradCalculator.StandardCalculator()
		case 2:
			evenOrOdd.EvenOrOdd()
		case 3:
			russianRoulette.RussianRoulette()
		case 4:
			fmt.Println("This option is in development")
		case 5:
			fmt.Println("The end of current session")
			return
		default:
			fmt.Println("This option is not available")
		}
	}
}
