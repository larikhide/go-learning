package quenec

import "fmt"

//cоздать очередь
var q []int

//AddQ добавить новый элемент в очередь
func AddQ(str string) {
	q = append(q, 1)
}

//DiscardQ удалить из очереди
func DiscardQ() {
	q = q[1:]
	if len(q) == 0 {
		fmt.Println("Очереди нет")
	}
}
