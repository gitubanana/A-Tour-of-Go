package main

import "fmt"

func main() {
	sum := 0
	for i := 0; i < 10; i++ { // 굳이 0부터 시작할 필요가 있나 싶다.
		sum += i
	}
	fmt.Println(sum)
}
