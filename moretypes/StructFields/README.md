### Struct Fields
```go
type Vertex struct {
  X int
  Y int
}

func main() {
  v := Vertex{1, 2}
  v.X = 4
  fmt.Println(v.X)
}
```
`struct`의 필드는 `.`을 이용해 접근할 수 있다.<br>

### go run
![image](https://github.com/user-attachments/assets/53915029-f3eb-4d7b-a761-d818c455e59a)


### Reference
[Moretypes/Struct Fields](https://go.dev/tour/moretypes/3)<br>
