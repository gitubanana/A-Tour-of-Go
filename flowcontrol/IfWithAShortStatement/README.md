### If with a short statement
```go
func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  }
  return lim
}
```
`Go`의 `if`문은 `for`문처럼 `초기화`를 할 수 있다.<br>
여기서 생긴 변수는 `if`문 안에서만 사용할 수 있다.<br>

### go run
![image](https://github.com/user-attachments/assets/200e2bfd-db74-4365-8151-4baae45cc149)

### Reference
[Flowcontrol/If with a short statement](https://go.dev/tour/flowcontrol/6)<br>
