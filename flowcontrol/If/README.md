### If
```go
func sqrt(x float64) string {
  if x < 0 {
    return sqrt(-x) + "i"
  }
  return fmt.Sprint(math.Sqrt(x)) // string으로 변환해서 반환
}
```
`Go`의 `if`문은 `for`이랑 똑같이 `()`는 필요없고, `{}`는 필요하다.<br>

### go run
![image](https://github.com/user-attachments/assets/2a1f40e7-5407-460e-8a75-b45870911755)

### Reference
[Flowcontrol/IF](https://go.dev/tour/flowcontrol/5)<br>
