package main

import (
	"fmt"
)

func main() {
	persons := map[string][]string{
		"you": []string{
			"Alice",
			"Bob",
			"Claire",
		},
		"Bob": []string{
			"Anuj",
			"Peggy",
		},
		"Alice": []string{
			"Peggy",
		},
		"Claire": []string{
			"Thom",
			"Jonny",
		},
		"Anuj":  []string{},
		"Peggy": []string{},
		"Thom":  []string{},
		"Jonny": []string{},
	}
	visited := []string{}
	bfs("Bob", persons, func(name string) { //в функцию bfs передаем аргументы: имя, кого проверяем и его друзей; исходную хеш-таблицу;
		visited = append(visited, name) // и функцию, которая к списку проверенных имен добавляет вновь проверенное имя
	})
	fmt.Println(visited)
}

/* функция поиска в ширину.
Аргументы: имя, того кого проверяем; исходная структура по которой обходим и
функция fn которая что-то делает с результатом поиска */
func bfs(start string, persons map[string][]string, fn func(string)) {
	frontier := []string{start}  //изначально проверяемое имя
	visited := map[string]bool{} //
	next := []string{}           //

	for 0 < len(frontier) { //пока у проверяемой строки есть длина, то продолжать расчет
		next = []string{}               //
		for _, name := range frontier { //обход по значениям строки
			visited[name] = true
			fn(name)
			for _, n := range bfsFrontier(name, persons, visited) {
				next = append(next, string(n))
			}
		}
		frontier = next
	}
}

/* функция берет имя, карту, есть ли значение под ключом, выдает массив строк */
func bfsFrontier(name string, persons map[string][]string, visited map[string]bool) []string {
	next := []string{}                                                      //объявляем массив строк
	iter := func(n string) bool { _, ok := visited[string(n)]; return !ok } //объявляем функцию как значение. Получаем значение из карты по имени. Возвращаем хуй знает что и зачем
	for _, n := range persons[name] {                                       //обходим исходную мапу по значениям
		if iter(n) { //если
			next = append(next, string(n))
		}
	}
	return next
}
