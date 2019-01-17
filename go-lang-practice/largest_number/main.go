package main

import "fmt"

func main() {
	fmt.Println("Enter 3 numbers :")
	var first, second, third int
	input(first, second, third)
}

func input(f int, s int, t int) {
	fmt.Scanln(&f)
	fmt.Scanln(&s)
	fmt.Scanln(&t)
	compare(f, s, t)
}

func compare(first int, second int, third int) {

	if first >= second && first >= third {
		fmt.Println(first, "is Largest among three numbers.")
	}
	if second >= first && second >= third {
		fmt.Println(second, "is Largest among three numbers.")
	}
	if third >= first && third >= second {
		fmt.Println(third, "is Largest among three numbers")
	}
	// if first == second && first == third {
	// 	fmt.Println("All Three are Equal Numbers!")
	// }
}
