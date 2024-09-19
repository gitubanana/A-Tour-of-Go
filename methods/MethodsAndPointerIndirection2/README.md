### Methods and pointer indirection 2
값을 받는 함수는 값을 받아야 한다.<br>
```go
var v Vertex
fmt.Println(AbsFunc(v)) // OK
fmt.Println(AbsFunc(&v)) // Compiler error!
```

반면에, `value receiver`를 받는 `method`는 값과 포인터 둘 다 받아도 된다.<br>
```go
var v Vertex
fmt.Println(v.Abs()) // OK

p := &v

// (*p).Abs()로 해석된다.
fmt.Println(p.Abs()) // OK
```

### go run
![image](https://github.com/user-attachments/assets/7799b9c7-75e7-4423-9484-a1fa73c59bf9)


### Reference
[Methods/Methods and pointer indirection](https://go.dev/tour/methods/7)<br>
