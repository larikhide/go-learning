package main

import "fmt"

//Machine is a structire
type Machine struct {
	Label   string
	Year    uint
	Volume  float64
	Engine  bool
	Windows bool
	Fill    float64
}

func main() {
	var lightCar = Machine{
		Label:   'lada',
		Year:    2007,
		Volume:  40.2,
		Engine:  true,
		Windows: true,
		Fill:    0.5,
	}
	fmt.Println(lightCar)
}
