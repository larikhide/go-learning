package main

import (
	"fmt"
	"math"
)

func main() {
	// 1st task
	var rubles float32
	fmt.Println("How much rubles do you want to exchange?")
	fmt.Scanf("%f", &rubles)
	fmt.Printf("It will be %f in dollars\n", currencyConverter(rubles))

	// 2nd task
	fmt.Println(triangleCalc(3, 4))

	//3rd task
	var rate, deposit float64

	fmt.Println("insert percent rate")
	fmt.Scanf("%f", &rate)

	fmt.Println("insert deposit amount")
	fmt.Scanf("%f", &deposit)

	for i := 1; i <= 5; i++ {
		deposit += deposit * rate / 100
	}

	fmt.Printf("your sum after 5 years is %f\n", deposit)
}

func currencyConverter(r float32) (d float32) {
	const rate = 63.1
	d = r / rate
	return
}

func triangleCalc(a, b float64) (hypo, perim, area float64) {
	hypo = math.Hypot(a, b)
	perim = a + b + hypo
	area = a * b / 2
	return
}
