### If and else
```go
func pow(x, n, lim float64) float64 {
  if v := math.Pow(x, n); v < lim {
    return v
  } else {
    fmt.Printf("%g >= %g\n", v, lim)
  }
  // can't use v here, though
  return lim
}
```
`if`문에서 초기화된 변수는 `else`에서도 사용할 수 있다.<br>

### go run
![image](https://github.com/user-attachments/assets/e0fbdd7d-5be0-43ba-a96d-cb2981decb9a)


### Reference
[Flowcontrol/If and else](https://go.dev/tour/flowcontrol/7)<br>
