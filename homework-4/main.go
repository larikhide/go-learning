package main

import "fmt"

// 1. написать свой интерфейс и несколько структур удволетворяющих ему
type getHomer interface {
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
	fmt.Printf("Вуки под именем %s живет в деревне %s в дупле под номером %d", w.Name, w.Village, w.Hollow)
}

func (h Human) home() {
	fmt.Println("Человеки живут в бетонных коробках")
}

func main() {
	var chuwi getHomer = Wookie{Name: "Choobaka", Village: "Hairybugs", Hollow: 3}
	var khan getHomer = Human{}
	chuwi.home()
	khan.home()
}
