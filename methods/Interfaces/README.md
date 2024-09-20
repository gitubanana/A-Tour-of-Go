### Interfaces
`interface type`은 `method`의 모음으로 정의될 수 있다.<br>
`interface type`의 값은 그 `method`를 구현해 놓은 어떠한 값이든 가질 수 있다.<br>
```go
type Abser interface {
  Abs() float64 // 구현해야 하는 method
}

type MyFloat float64
func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
return float64(f)
}

type Vertex struct {
  X, Y float64
}
func (v *Vertex) Abs() float64 {
  return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
  var a Abser
  
  f := MyFloat(-math.Sqrt2)
  v := Vertex{3, 4}
  
  a = f  // Myfloat implements Abser
  a = &v // *Vertex implements Abser
  
  // v is a Vertex -> Vertex doesn't have a method named Abs -> doesn't implement Abser.
  // a = v // Compile Error!
  
  fmt.Println(a.Abs())
}
```

### go run
![image](https://github.com/user-attachments/assets/72b506a4-e40a-4b91-8eef-1744b245f9ac)


### Reference
[Methods/Interfaces](https://go.dev/tour/methods/9)<br>
