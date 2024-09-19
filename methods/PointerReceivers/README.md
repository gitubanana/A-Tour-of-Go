### Pointer receivers
`pointer receiver`를 가지고 있는 `method`를 마들 수 있다.<br>
즉, `receiver`의 타입이 `*T`일 수 있다. (`T`는 `pointer`가 아닌 타입)<br>

`pointer receiver`를 통해서, `receiver`가 가리키는 값을 바꿀 수 있다.<br>
포인터가 아닌 그냥 `value receiver`라면, 원래 값의 복사본을 바꾸기 바꾼다.<br>

```go
type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v *Vertex) Scale(f float64) {
  v.X = v.X * f
  v.Y = v.Y * f
}

func main() {
  v := Vertex{3, 4}
  v.Scale(10)
  fmt.Println(v.Abs())
}
```

### go run
![image](https://github.com/user-attachments/assets/f971e8db-b917-41dd-9d06-bc6bbc9236e2)


### Reference
[Methods/Pointer Receivers](https://go.dev/tour/methods/4)<br>
