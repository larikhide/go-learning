package main

import (
	"fmt"
)

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
	fmt.Printf("Вуки под именем %s живет в деревне %s в дупле под номером %d\n", w.Name, w.Village, w.Hollow)
}

func (h Human) home() {
	fmt.Printf("Человек под именем %s живет на улице %s в доме под номером %d в квартире %d\n", h.Name, h.Street, h.Build, h.Apartment)
}

// 2. Создать тип, описывающий контакт в телефонной книге. Создать псевдоним типа телефонной книги (массив контактов) и реализовать для него интерфейс Sort{}.

// Contact создать тип описывающий контакт в телефонной книге
type Contact struct {
	Name    string
	Number  int32
	Address string
}

//ContactBook создать псевдоним типа телефонной книги (массив контактов)
type ContactBook []Contact

//Sort интерфейс который нужно удволетворить структурой, у которой есть метод SortingMethod
type Sort interface {
	SortingMethod()
}

//SortingMethod метод, который создан для удволетворения интерфейса Sort{}
func (cb ContactBook) SortingMethod() {
	fmt.Printf("Тут должнен быть написан алгоритм сортировки для книги контактов %v\n", cb)
}

func main() {
	chuwi := Wookie{Name: "Choobaka", Village: "Hairybugs", Hollow: 3}
	khan := Human{Name: "Khan Solo", Street: "Puskins", Apartment: 5}
	chuwi.home()
	khan.home()

	andrey := Contact{Name: "Andrusha", Number: 1245, Address: "BigHouse"}
	vladik := Contact{Name: "Vladushka", Number: 5421, Address: "SmallHouse"}
	book := make(ContactBook, 0)
	book = append(book, andrey, vladik)
	book.SortingMethod()
}
