package main

import (
	"fmt"

	"github.com/larikhide/GoSyntax/homework-4/calculator/pkg"
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

		res, err := calculator.Calculate(input)

		switch {
		case err == nil :
			fmt.Printf("Результат: %v\n", res)
		case input == "help":
			fmt.Println("Reference is here.\nFor exit inserts exit. SUDDENLY!")
		default:
			fmt.Println("Не удалось произвести вычисление")
		}
	}
}
