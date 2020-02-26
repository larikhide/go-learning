package main

import "fmt"

// 1. написать свой интерфейс и несколько структур удволетворяющих ему
type getHome interface {
	home()
}

//Wookie структура описывающая дом вуки
type Wookie struct {
	Name    string
	Village string
	Hollow  int
}

//Human структура описывающая дом человеков
type Human struct {
	Name      string
	Street    string
	Build     int
	Apartment int
}

//пытаюсь реализовать методы для каждой структуры, для того чтобы они удволетворяли интерфейсу

func (w Wookie) home() {
	fmt.Println("Вуки живут в деревне на деревьях")
}

func (h Human) home() {
	fmt.Println("Человеки живут в бетонных коробках")
}

func main() {
	var chuwi getHome = Wookie{}
	var khan getHome = Human{}
	chuwi.home()
	khan.home()
}
