package standaradCalculator

import (
	"fmt"
)

func StandardCalculator() {

	fmt.Println("\n This is standard calculator. You should type 3 elements: ")
	fmt.Println("1) first operand 2) operator 3) second operand ")
	fmt.Println(" available 5 operators:\n ( +  -  *  / % ) ")
	fmt.Println(" You have to use space between every element!")
	fmt.Println(" Usage ex:	2 * 2	6 - 4	5 + 5	11 / 2")

	var a, b float64
	var c string

	for {
		for {
			fmt.Println("\nPress 0 to comeback to main menu")
			fmt.Println("\n>>> Standard Calculator <<<\n")

			fmt.Scanln(&a, &c, &b)

			// scanner := bufio.NewScanner(os.Stdin)
			// scanner.Scan()
			// a, _ := strconv.ParseFloat(scanner.Text(), 64)

			// c := scanner.Text()
			// b, _ := strconv.ParseFloat(scanner.Text(), 64)

			// scanner3 := bufio.NewScanner(os.Stdin)
			// scanner3.Scan()
			// c := scanner.Text()

			// scanner2 := bufio.NewScanner(os.Stdin)
			// scanner2.Scan()
			// b, err2 := strconv.ParseFloat(scanner.Text(), 64)

			// if err1 != nil {
			// 	fmt.Println("First argument is invalid operator, try again!")
			// 	break
			// }

			// if err2 != nil {
			// 	fmt.Println("Second argument is invalid operator, try again!")
			// 	break
			// }

			if a == 0 {
				return
			}

			switch c {
			case "+":
				fmt.Printf("%.2f %s %.2f = %.2f", a, c, b, a+b)
			case "-":
				fmt.Printf("%.2f %s %.2f = %.2f", a, c, b, a-b)
			case "*":
				fmt.Printf("%.2f %s %.2f = %.2f", a, c, b, a*b)
			case "/":
				if b == 0 {
					fmt.Print("Math mistake! Division by zero!")
				} else {
					fmt.Printf("%.2f %s %.2f = %.2f", a, c, b, a/b)
				}
			case "%":
				fmt.Printf("%d %s %d = %d", int(a), c, int(b), int(a)%int(b))
			case "<":
				return
			default:
				fmt.Printf("You used %s,which is invalid operator\n", c)
			}
		}
		fmt.Println("")
	}
}
