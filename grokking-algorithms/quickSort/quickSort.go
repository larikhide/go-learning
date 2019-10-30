package main

import (
	"fmt"
)

/* func quickSort(ar []int32) {
	pivot := ar[len(ar)/2]
	left := 0
	right := len(ar) - 1
	if len(ar) < 2 {
		return
	}
	//проходимся по массиву
	//получаем индекс элемента, относительного которого слева - всё, что меньше
	//справа - больше
	for { //выполняется пока не произойдет return
		for ; ar[left] < pivot; left++ { //while ar[left] < pivot
		}
		for ; ar[right] > pivot; right-- { //same
		}
		if left >= right { //если условие выполнилось, то массив отсортирован на большее и меньшее относительно элемента под индексом right
			break
		}
		if ar[left] == ar[right] { //проверка на равенство элемента слева и справа
			break
		}
		ar[left], ar[right] = ar[right], ar[left]
	}
	quickSort(ar[:right])
	quickSort(ar[right:])
} */

func quickSort(ar []int32) {
	if len(ar) < 2 { //проверка на базовый случай
		return
	}
	pivot := ar[len(ar)/2]
	left := 0
	right := len(ar) - 1

Loop:
	for {
		switch {
		case ar[left] < pivot:
			left++
		case ar[right] > pivot:
			right--
		case (left >= right):
			break Loop
		case (ar[left] == ar[right]):
			left++
		}
		ar[left], ar[right] = ar[right], ar[left]
	}
	quickSort(ar[:right])
	quickSort(ar[right:])
}

func main() {
	ar := []int32{1, 4, 0, 1, 0, 1, -14, 1}
	//ввод с клавы
	/* var (
		n  int
		el int32
	)
	fmt.Scan(&n)
	ar := make([]int32, n)
	for i := range ar {
		fmt.Scan(&el)
		ar[i] = el
	} */
	quickSort(ar)
	fmt.Print(ar)
}
