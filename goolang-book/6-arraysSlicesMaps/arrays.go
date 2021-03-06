package main

import "fmt"

func main() {
	var x [5]int
	x[4] = 100        // присвоить пятому элементу массива значение = 100. Массивы нумеруются с нуля
	fmt.Println(x)    // можно вывести так весь массив
	fmt.Println(x[4]) // а можно так пятый элемент

	// пример программы с использованием массива
	var y [5]float64 // создается массив длиной в 5 элементов
	y[0] = 98        // чтобы выделить следующее вхождение ctrl+d
	y[1] = 93        // чтобы выделить все вхождения ctrl+f2
	y[2] = 77        // чтобы выделить колонку, нужно зажать Alt+shift+(drag mouse)
	y[3] = 82        // чтобы переместить строку, нужно зажать alt и стрелки вверх-вниз
	y[4] = 83        // чтобы копировать строку, нужно зажать alt+shift и стрелки вверх-вниз

	var total float64 = 0 // подсчет общего кол-ва баллов
	for i := 0; i < 5; i++ {
		total += y[i]
	}
	fmt.Println(total / 5) // делим общую сумму баллов на количество элементов, чтобы узнать средний балл

	// оптимизация кода приведенного выше
	var q [5]float64
	q[0] = 98
	q[1] = 93
	q[2] = 77
	q[3] = 82
	q[4] = 83
	var totalOptimize float64 = 0
	for o := 0; o < len(q); o++ { // если надо будет изменить длину массива, то лучше оперировать len(q)
		totalOptimize += q[o]
	}
	fmt.Println(totalOptimize / float64(len(q))) // len (q) и тотал имеют разный тип. Конвертируем len(q) во float64


	// ещё оптимитизации

	// оптимизация кода приведенного выше
	var e = [5]float64 {98,93,77,82,83,} // короткая запись массивов. Можно в записывать в несколько строк, но после последнего элемента обязатоельно ставить ","
	var totalOptimize2 float64 = 0
	for _, value := range e { // В этом цикле i представляет текущую позицию в массиве, а value будет тем же самым что и x[i]. Мы использовали ключевое слово range перед переменной, по которой мы хотим пройтись циклом."_" вместо индекса элемента - сказать Go что эту переменную не юзаем
		totalOptimize2 += value
	}
	fmt.Println(totalOptimize2 / float64(len(e)))
}
