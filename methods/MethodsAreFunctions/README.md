### Methods are functions
`method`는 `receiver`란 인자를 가지고 있는 "함수"임을 기억하시라네요~!<br>
아래 `Abs`함수는 저번 함수와 기능이 정확히 똑같은 평범한 함수입니다.<br>
```go
type Vertex struct {
  X, Y float64
}

func Abs(v Vertex) float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3, 4}
  fmt.Println(Abs(v))
}
```

### go run
![image](https://github.com/user-attachments/assets/8eea423d-2ad8-4f0c-baff-2d394499c419)


### Reference
[Methods/Methods are functions](https://go.dev/tour/methods/2)<br>
