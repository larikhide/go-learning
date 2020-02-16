package main

import (
	"fmt"
	"math/big"
)

// 1. Написать функцию, которая определяет, четное ли число.
func evenNum(a int) {
	if a%2 == 0 {
		fmt.Printf("Число %v является четным числом\n", a)
	} else {
		fmt.Printf("Число %v не является четным числом\n", a)
	}
}

// 2. Написать функцию, которая определяет, делится ли число без остатка на 3.
func multOfThree(a int) {
	if a%3 == 0 {
		fmt.Printf("Число %v делится на 3 без остатка\n", a)
	} else {
		fmt.Printf("Число %v не делится на 3 без остатка\n", a)
	}
}

// 3. Написать функцию, которая последовательно выводит на экран 100 первых чисел Фибоначчи, начиная с 0.
func fibHundred(a, b *big.Int) {
	fmt.Printf("Первые 100 чисел Фибоначчи: ")
	for i := 0; i <= 98; i++ {
		fmt.Printf("%v ", a)
		// складывает а и Б , результат записывает в а
		a.Add(a, b)
		a, b = b, a
	}
	return
}

func sieveOfEratos(n int) {
	//создать булевый массив зполненный false
	nums := make([]bool, n)

	//вычеркнуть числа от 2p, 3p .. до n
	for i := 2; i*i < n; i++ {
		if nums[i] == false {
			for j := i * i; j < n; j += i {
				nums[j] = true
			}
		}
	}

	//вывести на экран натуральные числа
	fmt.Printf("\nПростые числа меньше %v: ", n)
	for i := 2; i < n; i++ {
		if nums[i] == false {
			fmt.Printf("%v ", i)
		}
	}
}

func main() {
	a := 6
	evenNum(a)
	multOfThree(a)
	fibHundred(big.NewInt(0), big.NewInt(1))
	sieveOfEratos(100)
}
