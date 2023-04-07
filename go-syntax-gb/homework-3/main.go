package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"

	"github.com/larikhide/go-learning/go-syntax-gb/homework-3/quenec"
)

// Machine 1. Описать несколько структур
type Machine struct {
	Label   string
	Year    uint
	Volume  float64
	Engine  bool
	Windows bool
	Fill    float64
}

func main() {
	//2. инициализировать несколько структур
	var lightCar = Machine{
		Label:   "lada",
		Year:    2007,
		Volume:  40.2,
		Engine:  true,
		Windows: true,
		Fill:    0.5,
	}

	var heavyCar = Machine{
		Label:   "Scania",
		Year:    2014,
		Engine:  false,
		Windows: true,
		Fill:    1,
	}

	//2. Применить к ним различные действия.
	lightCar.Year++
	heavyCar.Windows = false
	heavyCar.Volume = lightCar.Volume + 1050

	//2. Вывести значения свойств экземпляров в консоль.
	fmt.Println(lightCar, heavyCar)
	fmt.Println(heavyCar.Label, lightCar.Year)

	//3. Реализовать очередь fifo
	quenec.AddQ("Первый")
	quenec.AddQ("Второй")
	quenec.AddQ("Третий")

	fmt.Println(quenec.DiscardQ())
	fmt.Println(quenec.DiscardQ())

	quenec.AddQ("Четвертый, который по идее уже второй")

	fmt.Println(quenec.DiscardQ())
	fmt.Println(quenec.DiscardQ())
	fmt.Println(quenec.DiscardQ())

	//* 4. Внести в телефонный справочник дополнительные данные. Реализовать сохранение json-файла на диске с помощью пакета ioutil при изменении данных.
	//cоздали книгу, записали двоих
	addressBook := make(map[string]int)
	addressBook["Masha"] = 8918745875
	addressBook["Vasiliy"] = 8916547155

	// из книги сделали массив байт
	aB, err := json.Marshal(addressBook)
	if err != nil {
		log.Println(err)
		return
	}

	// записали массив байт в файл
	err = ioutil.WriteFile("myAddressBook", aB, 0644)
	if err != nil {
		log.Fatal(err)
	}

	//прочитали массив байт из файла
	content, err := ioutil.ReadFile("myAddressBook")
	if err != nil {
		log.Fatal(err)
	}

	//массив байт в карту
	err = json.Unmarshal(content, &aB)

	//добавили нового человека
	addressBook["Ahmed"] = 8922888294

	// из книги сделали массив байт
	aB, err = json.Marshal(addressBook)
	if err != nil {
		log.Println(err)
		return
	}

	// записали массив байт в файл
	err = ioutil.WriteFile("myAddressBook", aB, 0644)
	if err != nil {
		log.Fatal(err)
	}
}
