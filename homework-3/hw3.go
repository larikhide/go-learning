package main

import "fmt"

//Machine 1. Описать несколько структур
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
}
