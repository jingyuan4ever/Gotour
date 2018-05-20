package main

import "golang.org/x/tour/tree"
import "fmt"

// Walk 步进 tree t 将所有的值从 tree 发送到 channel ch。
func Walk(t *tree.Tree, ch chan int) {
	if t.Left != nil {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if t.Right != nil {
		Walk(t.Right, ch)
	}
}

// Same 检测树 t1 和 t2 是否含有相同的值。
func Same(t1, t2 *tree.Tree) bool {
	ch1 := make(chan int)
	defer close(ch1)
	ch2 := make(chan int)
	defer close(ch2)

	go Walk(t1, ch1)
	go Walk(t2, ch2)
	b := true
	for i := 0; i < 10; i++ {
		v1 := <-ch1
		v2 := <-ch2
		if v1 != v2 {
			b = false
		}
	}
	return b
}

func main() {
	fmt.Println(Same(tree.New(1), tree.New(2)))
	fmt.Println(Same(tree.New(3), tree.New(3)))
}
