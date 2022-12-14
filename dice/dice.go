// Program to demonstrate triangle pic that you create with paticular parameters

package dice

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"time"
)

func Dice() {

	fmt.Println("\nType ANY NUMBER to shoot the dice!!!")

	for {
		for {

			fmt.Println("\nPress 0 to comeback to main menu")
			fmt.Println("\n>>> Russian Roulette <<<\n")

			scanner := bufio.NewScanner(os.Stdin)
			scanner.Scan()
			input, err := strconv.ParseInt(scanner.Text(), 10, 64)
			if err != nil {
				fmt.Println("Again, please, type any number ")
				fmt.Println("to shoot the dice (except 0, of course)")
				break
			}
			rand.Seed(time.Now().UnixNano())
			dice := 6
			r := rand.Intn(dice)

			if input == 0 {
				return
			} else {
				switch r {
				case 1:
					fmt.Println("You've got  1\n ")
					break
				case 2:
					fmt.Println("You've got  2\n ")
					break
				case 3:
					fmt.Println("You've got  3\n ")
					break
				case 4:
					fmt.Println("You've got  4\n ")
					break
				case 5:
					fmt.Println("You've got  5\n ")
					break
				case 0:
					fmt.Println("You've got  6\n")
					break

				}
			}
		}
	}
}
