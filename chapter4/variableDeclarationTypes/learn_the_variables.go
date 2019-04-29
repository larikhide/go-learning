package main

import "fmt"

func main() {
	var q string = "Hello World"
	fmt.Println(q) /* можно сделать так */

	var w string
	w = "Hello World"
	fmt.Println(w) /* а можно так */

	var e string
	e = "first"
	fmt.Println(e)
	e = "second"
	fmt.Println(e) /* можно так */

	var r string
	r = "first"
	fmt.Println(r)
	r = r + "second"
	fmt.Println(r) /* но так пизже? */

	// сравнивание
	var t string = "hello"
	var y string = "world"
	fmt.Println(t == y)

	var u string = "hello"
	var i string = "hello"
	fmt.Println(u == i)

	// задание переменных без объявления переменной и ее типа
	o := "bryaa" /* можно задать переменную без указания типа и то что это переменная */

	var p = "Hello World" /* можно объявить переменную без объявления типа */

	a := 5 /* КРУЧЕ ВСЕГО когда мега кратко */
	fmt.Println(o, p, a)

	// как правильно назвать переменную:
	// 1) должны начинаться с буквы
	// 2) содержать буквы, цифры и знак "_"
	s := "Max"
	fmt.Println("My dog's name is", s)

	name := "Max"
	fmt.Println("My dog's name is", name)
	dogsName := "Max" /* так называемый camelBack, первая буква первого слова записывается в нижнем регистре, первая буква последующих слов записывается в верхнем регистре. ТАК КРУЧЕ ВСЕГО */
	fmt.Println("My dog's name is", dogsName)
}
