### Methods and pointer indirection
포인터를 인자로 받는 함수는 무조건 포인터를 받아야 한다.<br>
```go
var v Vertex
ScaleFunc(v, 5) // Compile error!
ScaleFunc(&v, 5) // OK
```

`pointer receiver`를 가지고 있는 `method`는 포인터와 값 둘 다 받아도 된다.<br>
```go
var v Vertex

// Scale 메소드가 pointer receiver를 받기 때문에
// Go에서 편의를 위해 (&v).Scale(5)로 해석을 한다.
v.Scale(5) // OK

p := &v
p.Scale(10) // OK
```

### go run
![image](https://github.com/user-attachments/assets/3a1eba0a-a7f2-40de-8213-b4bd7faccfa1)


### Reference
[Methods/Methods and pointer indirection](https://go.dev/tour/methods/6)<br>
