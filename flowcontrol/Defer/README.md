### Defer
```go
func main() {
  defer fmt.Println("world")

  fmt.Println("hello")
}
```
`defer`문은 현재 함수가 끝난 후에 함수를 부른다. (끝날 때까지 함수의 실행을 미룬다.)<br>
`defer`함수가 부르는 함수의 인자는 바로 계산되지만, 함수는 감싸고 있는 함수의 `return`후에야 실행된다.<br>
(정말 신기한 기능이다.)<br>

### go run
![image](https://github.com/user-attachments/assets/c5ab4ef3-3c0d-4df4-ae75-51597ea1209b)

### Reference
[Flowcontrol/Defer](https://go.dev/tour/flowcontrol/12)<br>
