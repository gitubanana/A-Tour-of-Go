## math/rand
### rand.Intn
```go
func Intn(n int) int
```
랜덤한 정수 x (0 <= x < n)를 리턴한다.
(n <= 0이면, 런타임에서 panic)

## fmt
### fmt.Println
```go
func Println(a ...any) (n int, err error)
```
Standard output으로 피연산자 출력. (피연산자 사이에는 공백 출력, 마지막에는 개행 출력)
출력한 바이트 수와 출력 에러를 리턴한다.

## Reference
[rand.Intn](https://pkg.go.dev/math/rand#Intn)<br>
[fmt.Println](https://pkg.go.dev/fmt#Println)<br>
