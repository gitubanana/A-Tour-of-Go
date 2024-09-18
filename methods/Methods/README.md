### Methods
`Go`에는 `class`가 없다.<br>
하지만, 타입의 `method`를 정의할 순 있다.<br>

`method`는 `receiver`이라는 특별한 인자를 가지고 있는 함수이다.<br>
`receiver`는 `func`키워드와 함수명 사이에 쓰면 된다.<br>
```go
type Vertex struct {
  X, Y float64
}

func (v Vertex) Abs() float64 { // v라는 Vertex 타입의 receiver
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  v := Vertex{3, 4}
  fmt.Println(v.Abs())
}
```

### go run
![image](https://github.com/user-attachments/assets/e1ddc930-63ed-484b-8da2-ab4480d31c5e)


### Reference
[Methods/Methods](https://go.dev/tour/methods/1)<br>
