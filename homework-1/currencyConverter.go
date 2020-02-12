package main

import "fmt"

const rate = 63.1

func main() {
	var rubles float32
	fmt.Println("How much rubles do you want to exchange?")
	fmt.Scanf("%f", &rubles)
	dollars := rubles / rate
	fmt.Printf("It will be %f in dollars\n", dollars)
}
