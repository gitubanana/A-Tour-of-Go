package main

import "fmt"

type Vertex struct {
	X, Y int
}

var (
	v1 = Vertex{1, 2}  // Vertex 타입
	v2 = Vertex{X: 1}  // Y:0은 생략됨
	v3 = Vertex{}      // X:0과 Y:0 생략됨
	p  = &Vertex{1, 2} // *Vertex 타입
)

func main() {
	fmt.Println(v1, p, v2, v3)
}
