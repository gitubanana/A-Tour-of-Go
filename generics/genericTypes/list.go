package main

import "fmt"

type List[T any] struct {
	next *List[T]
	val  T
}

func (l *List[T]) printAll() {
	cur := l
	for ; cur != nil; cur = cur.next {
		fmt.Printf("%v -> ", cur.val)
	}

	fmt.Printf("nil\n")
}

func (l *List[T]) addBack(val T) {
	cur := l
	for ; cur.next != nil; cur = cur.next {
	}

	cur.next = new(List[T])
	cur.next.val = val
}

func main() {
	words := []string{"like", "euiclee", "!"}
	root := List[string]{val: "I"}

	for _, word := range words {
		root.addBack(word)
	}
	root.printAll()
}
