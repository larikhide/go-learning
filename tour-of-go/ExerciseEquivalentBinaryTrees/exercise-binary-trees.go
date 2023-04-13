package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value

	if t.Right != nil {
		Walk(t.Right, ch)
	}

}

func WalkAndClose(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch)
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	ch2 := make(chan int)
	go WalkAndClose(t1, ch1)
	go WalkAndClose(t2, ch2)

	return <-ch1 == <-ch2
}

func main() {
	//ch := make(chan int)
	/* go WalkAndClose(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	} */

	fmt.Println(Same(tree.New(2), tree.New(2)))
	fmt.Println(Same(tree.New(2), tree.New(3)))
	fmt.Printf("New 2 is %v \n New 3 is %v", tree.New(2), tree.New(2))
}
