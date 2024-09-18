### Methods continued
`non-struct` 타입에도 `method`를 정의할 수 있다.<br>
현재 패키지에 정의되어 있는 타입의 `method`만 정의할 수 있다.<br>
따라서, 다른 패키지에 정의되어 있는 타입의 `method`는 정의할 수 없다.<br>
(built-in 타입인 `int`도 마찬가지이다.)<br>
```go
type MyFloat float64 // float64의 method는 정의할 수 없기 때문에, 새로운 타입을 만들었다.

func (f MyFloat) Abs() float64 {
  if f < 0 {
    return float64(-f)
  }
  return float64(f)
}
```

### go run
![image](https://github.com/user-attachments/assets/f543f5e1-209c-4ea2-8a3c-0a421f9bf276)


### Reference
[Methods/Methods continued](https://go.dev/tour/methods/3)<br>
