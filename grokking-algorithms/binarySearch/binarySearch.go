package main

import "fmt"

// функция возвращает позицию загаданного числа
func binarySearch(list []int, item int) (id int) {
	mid := int(float64(len(list)) / 2.0) //так нормально делать?
	switch {
	case len(list) == 0:
		id = -1 //нет элементов, нет числа, разобраться в чем проблема когда вводишь ряд без искомого числа
	case list[mid] == item:
		id = mid
	case list[mid] > item:
		id = binarySearch(list[:mid], item)
	case list[mid] < item:
		id = binarySearch(list[mid+1:], item)
		id += mid + 1
	}
	return
}

func main() {
	list := make([]int, 5)
	for i := 0; i < 5; i++ {
		fmt.Scan(&list[i])
	}
	fmt.Println(binarySearch(list, 4))
}
