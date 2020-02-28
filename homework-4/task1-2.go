package main

import (
	"fmt"
	"sort"
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

//ContactBook создать псевдоним типа телефонной книги (массив нтактов)
type ContactBook []Contact

//методы которые позволят массиву контактов удволетворить интерфейс Sort{}
func (cb ContactBook) Len() int           { return len(cb) }
func (cb ContactBook) Swap(i, j int)      { cb[i], cb[j] = cb[j], cb[i] }
func (cb ContactBook) Less(i, j int) bool { return cb[i].Name < cb[j].Name }

func main() {
	//использую функции для первого задания
	chuwi := Wookie{Name: "Choobaka", Village: "Hairybugs", Hollow: 3}
	khan := Human{Name: "Khan Solo", Street: "Puskins", Apartment: 5}
	chuwi.home()
	khan.home()

	// использую фунции для второго задания
	book := []Contact{
		{"Basya", 9999, "Pushkina Kolotushkina"},
		{"Andrey", 7777, "Makarova 45"},
		{"Zlata", 5555, "Vilkina Pupirkina"},
	}
	fmt.Println(book)
	sort.Sort(ContactBook(book))
	fmt.Println(book)
}
