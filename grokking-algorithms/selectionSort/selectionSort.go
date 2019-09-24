package main

import (
	"fmt"
)

/* // функция нахождения минимального числа в слайсе
func findSmallest(list []int) int {
	smallest := list[0]
	//smallestID := list[0],
	for _, min := range list {
		if min < smallest {
			smallest = min
			//smallestID = i
		}
	}
	return smallest
}

//функция сортировки выбором после нахождения минимального
func selectionSort(list []int) (sortedList []int) {
	for _, el := range list {
		el = findSmallest(list)
		sortedList = append(sortedList, el)
	}
	return
} */

/* //функция сортировки выбором, всё в одной. не закончено
func selectionSort(list []int) (sortedList []int) {
	smallest := list[0]
	for _, min := range list {
		if min < smallest {
			smallest = min
		}
		sortedList = append(sortedList, smallest)
	}
	return
} */

func selectionSort(list []int) []int {
	for i := 0; i < len(list)-1; i++ {
		id, min := i, list[i]
		for j := i + 1; j < len(list); j++ {
			if list[j] < min {
				id, min = j, list[j]
			}
		}
		list[id], list[i] = list[i], min
	}
	return list
}

func main() {
	list := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&list[i])
	}
	fmt.Println(selectionSort(list))
}
