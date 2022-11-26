// the Program to check for a number even or odd

package evenOrOdd

import (
	"fmt"
)

func EvenOrOdd() {

	var a int

	for {

		fmt.Println("\nPress 00 to comeback to main menu")
		fmt.Println("\n>>> Even or Odd checker <<<\n")

		fmt.Scanln(&a)

		// if a != int(a) {
		// 	fmt.Println("Incorrect input")
		// 	fmt.Println("You should type integer number")
		// }
		if a == 00 {
			return
		}
		switch a % 2 {
		case 0:
			fmt.Printf("%d is even\n", a)
		case 1, -1:
			fmt.Printf("%d is odd\n", a)
		}
	}
}
