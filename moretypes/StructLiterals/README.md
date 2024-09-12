### Struct literals
```go
var (
  v1 = Vertex{1, 2}  // has type Vertex
  v2 = Vertex{X: 1}  // Y:0 is implicit
  v3 = Vertex{}      // X:0 and Y:0
  p  = &Vertex{1, 2} // has type *Vertex
)
```
`struct literals`란 필드를 초기화하면서 새로 생성된 `struct`를 의미한다.<br>
`Name:`을 통해서, 초깃값을 나열할 수 있다.<br>
(나열하는 순서는 상관없다.)

### go run
![image](https://github.com/user-attachments/assets/7988ff51-ff5d-486a-b53e-7ec6a36fbfa3)


### Reference
[Moretypes/Struct literals](https://go.dev/tour/moretypes/5)<br>
