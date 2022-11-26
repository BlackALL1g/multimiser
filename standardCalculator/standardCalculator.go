package standaradCalculator

import "fmt"

func StandardCalculator() {

	var a, b float64
	var c string

	fmt.Println("\n This is standard calculator. You should type 3 elements: ")
	fmt.Println("1) first operand 2) operator 3) second operand ")
	fmt.Println(" available 5 operators:\n ( +  -  *  / % ) ")
	fmt.Println(" You have to use space between every element!")
	fmt.Println(" Usage ex:	2 * 2	6 - 4	5 + 5	11 / 2")

	for {
		fmt.Println("\nPress 00 to comeback to main menu")
		fmt.Println("\n>>> Standard Calculator <<<\n")

		fmt.Scanln(&a, &c, &b)

		if a == 00 {
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
		fmt.Println("")
	}
}
