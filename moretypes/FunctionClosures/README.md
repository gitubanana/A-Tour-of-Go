### Function closures
```go
func adder() func(int) int {
  sum := 0
  return func(x int) int {
    sum += x
    return x
  }
}

func main() {
  pos, neg := adder(), adder()
  for i := 0; i < 10; i++ {
    fmt.Println(
      pos(i),
      neg(-2*i),
    )
  }
}
```
`closure`란 자기 밖에 있는 변수를 참조하는 함수값이다.<br>
`closure`는 그 변수를 접근하고, 대입할 수 있다.<br>
(그 함수는 변수에 묶여있다.)<br>

위 예에서는, `adder`함수는 `closure`를 리턴한다.<br>
각 `closure`는 각 `sum` 변수에 묶여있다.<br>


### go run
![image](https://github.com/user-attachments/assets/9db7794c-7b8d-47c5-bed8-31a53b0c6227)


### Reference
[Moretypes/Function closures](https://go.dev/tour/moretypes/25)<br>
