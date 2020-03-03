package main

import (
	"fmt"
	"time"
)

//1/ Уберите из первого примера (Фибоначчи и спиннер) функцию, вычисляющую числа Фибоначчи. Как теперь заставить спиннер вращаться в течение некоторого времени? 10 секунд?
func spinner(delay time.Duration) {
	for {
		for _, r := range "-\\|/" {
			fmt.Printf("%c\r", r)
			time.Sleep(delay)
		}
	}
}

func main() {
	go spinner(50 * time.Millisecond)
	time.Sleep(1 * time.Second)
	// убирает кусок спиннера из вывода
	fmt.Printf("\r ")
}
