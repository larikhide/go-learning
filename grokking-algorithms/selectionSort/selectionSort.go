package main

import (
	"fmt"
)

// функция нахождения минимального числа в слайсе
func findSmallest(list []int) (smallest int) {
	smallest = list[0]
	for _, min := range list {
		if min < smallest {
			smallest = min
		}
	}
	return smallest
}

//функция сортировки выбором
func selectionSort(
//хуряить тут
)

func main() {
	list := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&list[i])
	}
	fmt.Println(findSmallest(list))
}
