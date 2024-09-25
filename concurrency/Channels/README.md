### Channels
`Channel`은 타입이 있는 전달자이다.<br>
`Channel`연산자 `<-`로 값을 주고 받을 수 있다.<br>
(화살표 방량으로 데이터가 옮겨진다.)
```go
ch <- v // Send v to channel ch.
v := <-ch // Receive from ch, and assign value to v.
```

`map`과 `slice`처럼, `channel`도 `make`로 만들어야 사용할 수 있다.<br>
```go
ch := make(chan int)
```

기본적으로, 값을 `받기`와 `보내기`는 다른 한 쪽이 준비될 때까지 블락된다.<br>
이 덕분에 `lock`혹은 상태 변수 없이도 `goroutine`을 동기화시킬 수 있다.<br>
```go
package main

import "fmt"

func sum(s []int, c chan int) {
	sum := 0
	for _, v := range s {
		sum += v
	}
	c <- sum // send sum to c
}

func main() {
	s := []int{7, 2, 8, -9, 4, 0}

	c := make(chan int)
	go sum(s[:len(s)/2], c)
	go sum(s[len(s)/2:], c)
	x, y := <-c, <-c // receive from c

	fmt.Println(x, y, x+y)
}

```

### go run
![image](https://github.com/user-attachments/assets/51403c0e-3f24-4830-98ef-77b983afc2f0)


### Reference
[Concurrency/Channels](https://go.dev/tour/concurrency/2)<br>
