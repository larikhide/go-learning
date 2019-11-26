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
	bfs("you", persons, func(name string) {
		visited = append(visited, name)
	})
	fmt.Println(visited)
}

func bfs(start string, persons map[string][]string, fn func(string)) {
	frontier := []string{start}
	visited := map[string]bool{}
	next := []string{}

	for 0 < len(frontier) {
		next = []string{}
		for _, name := range frontier {
			visited[name] = true
			fn(name)
			for _, n := range bfsFrontier(name, persons, visited) {
				next = append(next, string(n))
			}
		}
		frontier = next
	}
}

func bfsFrontier(name string, persons map[string][]string, visited map[string]bool) []string {
	next := []string{}
	iter := func(n string) bool { _, ok := visited[string(n)]; return !ok }
	for _, n := range persons[name] {
		if iter(n) {
			next = append(next, string(n))
		}
	}
	return next
}
