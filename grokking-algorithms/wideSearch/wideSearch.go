package main

import "fmt"

func main() {
	var nodes = map[int][]int{
		1:  []int{2, 3, 4},
		2:  []int{1, 5, 6},
		3:  []int{1},
		4:  []int{1, 7, 8},
		5:  []int{2, 9, 10},
		6:  []int{2},
		7:  []int{4, 11, 12},
		8:  []int{4},
		9:  []int{5},
		10: []int{5},
		11: []int{7},
		12: []int{7},
	}
	traders := []int{} //массив для чуваков, удволетворяющих условиям
	bfs(2, nodes, func(node int) {
		if node > 10 { //если чувак больше 10, то он подходит
			traders = append(traders, node)
		}
	})
	fmt.Println(traders)
}

//функция смотрит проверялся ли node из nodes и добавляет непроверенных в очередь
func queueMaker(node int, nodes map[int][]int, visited map[int]bool) []int {
	queue := []int{}                                               //массив в который будут добавляться непроверенные
	cheker := func(n int) bool { _, ok := visited[n]; return !ok } //проверка на существование элемента в visited. Если он уже там есть, то возвращает false, т.е. этот элемент в последствии не добавляется в очередь
	for _, n := range nodes[node] {
		if cheker(n) {
			queue = append(queue, n)
		}
	}
	return queue
}

func bfs(start int, nodes map[int][]int, fn func(int)) {
	first := []int{start}     //проверяемая очередь и первый в ней кандидат
	visited := map[int]bool{} //создание списка уже проверенных кандидатов
	next := []int{}           // массив для хранения нового состояния очереди

	for len(first) > 0 { //выполняется пока в first есть элементы
		next = []int{}               //точно не уверен, но по идее новое состояние очереди обнуляется, чтобы избежать бесконечного цикла, в котором мы раз за разом будем добалять ещё одну очередь по которой прошлись к концу цикла и постоянно бы проверяли одних и тех же.
		for _, node := range first { //обход друзей первого кандидата
			visited[node] = true                                 //помечается, что данный друг проверен
			fn(node)                                             //функция котороая что-то делает с другом
			for _, n := range queueMaker(node, nodes, visited) { //точно не уверен, но по идее проходимся по новому состоянию очереди, без только что просмотренного товарища
				next = append(next, n) //добавляем к кандидатам новое состояние очереди, после удаления из нее только что проверенного
			}
		}
		first = next //новое состояние очереди после проходки
	}
}
