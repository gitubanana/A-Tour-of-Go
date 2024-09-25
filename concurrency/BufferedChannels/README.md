### Buffered Channels
`channel`은 `buffered`될 수 있다.<br>
버퍼의 길이를 `make`의 두 번째 인자로 주면 된다.<br>
```go
ch := make(chan int, 100) 
```

`Buffered channel`에 **값 쓰기**는 `buffer`가 꽉 찼을 때 `Block`되고,<br>
`Buffered channel`에 **값 받기**는 `buffer`가 비었을 때 `Block`된다.<br>
```go
package main

import "fmt"

func main() {
  ch := make(chan int, 1)
  ch <- 1
  ch <- 2 // buffer가 꽉 찼을 때 값 쓰기를 해 Block되고,
          // 이 Block을 풀어줄 다른 쓰레드가 없기 때문에 무한정 기다리는 상태인 deadlock이 발생한다.
  /*
  fatal error: all goroutines are asleep - deadlock!
  
  goroutine 1 [chan send]:
  */

  fmt.Println(<-ch)
  fmt.Println(<-ch)
}
```

### go run
![image](https://github.com/user-attachments/assets/edd50c36-58b2-4979-9d71-6eacb22036b7)


### Reference
[Concurrency/Buffered Channels](https://go.dev/tour/concurrency/3)<br>
