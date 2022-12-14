// the Program to check for a number even or odd

package evenOrOdd

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func EvenOrOdd() {

	fmt.Println("\nPut any integer number and I will return the answer")
	for {
		for {

			fmt.Println("\nPress 0 to comeback to main menu")
			fmt.Println("\n>>> Even or Odd checker <<<\n")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			// if sc == "00" {
			// 	return
			// }
			input, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Println("It's not an int! type integer number")
				break
			}

			if input == 0 {
				return
			}

			switch input % 2 {
			case 0:
				fmt.Printf("%d is even\n", input)
			case 1, -1:
				fmt.Printf("%d is odd\n", input)
			}
		}
	}
}
