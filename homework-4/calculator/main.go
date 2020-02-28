package main

import (
	"fmt"
<<<<<<< HEAD

	"github.com/larikhide/GoSyntax/homework-4/calculator/pkg"
=======
	"github.com/larikhide/GoSyntax/homework-4/calculator"
>>>>>>> 29acc55fb251dfc75158db298029636c3838bc76
)

func main() {
	input := ""
	for {
		fmt.Print("> ")
		if _, err := fmt.Scanln(&input); err != nil {
			fmt.Println(err)
			continue
		}

		if input == "exit" {
			break
		}

<<<<<<< HEAD
=======
		if input == "help" {
			fmt.Println("тут типа текст справки")
		}

>>>>>>> 29acc55fb251dfc75158db298029636c3838bc76
		if res, err := calculator.Calculate(input); err == nil {
			fmt.Printf("Результат: %v\n", res)
		} else {
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
