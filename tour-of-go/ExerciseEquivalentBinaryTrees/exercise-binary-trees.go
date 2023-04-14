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
	//TODO: add test for WalkandCloses
	//ch := make(chan int)
	/* go WalkAndClose(tree.New(1), ch)
	for i := range ch {
		fmt.Println(i)
	} */
	//TODO: move test functions in a test package
	// type testCase struct {exp bool, rec bool}
	// var testCases []testCase {}
	fmt.Print("tree.New(1) == tree.New(1): ")
	expected := true
	received := Same(tree.New(1), tree.New(1))
	if received {
		fmt.Println("PASSED")
	} else {
		fmt.Printf("FAILED:\nneed %v\n got: %v", expected, received)
	}
	fmt.Print("tree.New(1) != tree.New(2): ")
	expected = false
	received = Same(tree.New(1), tree.New(2))
	if !received {
		fmt.Println("PASSED")
	} else {
		fmt.Printf("FAILED:\nneed %v\ngot: %v\n", expected, received)
	}
}
